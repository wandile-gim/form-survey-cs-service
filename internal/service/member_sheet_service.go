package service

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/api/sheets/v4"
)

var eventId = 2025524

type MemberSheetService struct {
	sheetReader *domain.SheetReader
	tracker     repository.TrackerRepository
	task        repository.TaskRepository
	smsService  Notifier
}

func NewMemberSheetService(s *domain.SheetReader, taskrepo repository.TaskRepository, tracker repository.TrackerRepository, notifier Notifier) MemberSheetService {
	return MemberSheetService{
		s, tracker, taskrepo, notifier,
	}
}

func (s MemberSheetService) convertDate(row interface{}) (time.Time, error) {
	var registeredAt time.Time
	if strings.Contains(row.(string), "오후") {
		row = strings.Replace(row.(string), "오후", "PM", 1)
	} else if strings.Contains(row.(string), "오전") {
		row = strings.Replace(row.(string), "오전", "AM", 1)
	}
	// 2024. 10. 11 오후 6:18:34
	if registeredAt2, err := time.Parse("2006. 1. 2 PM 3:04:05", row.(string)); err == nil {
		registeredAt = registeredAt2
	} else {
		log.Error().Msgf("datetime을 파싱하는데 실패했습니다. %v", err)
		return time.Time{}, err
	}
	return registeredAt, nil
}

func (s MemberSheetService) needToSkip(registeredAt time.Time, startIdx time.Time) bool {
	// registeredAt을 kst로 변경
	if registeredAt.Before(startIdx) || registeredAt.Equal(startIdx) {
		//log.Info().Msgf("이전 데이터 스킵: %v", registeredAt)
		//log.Info().Msgf("최신 스킵 데이터: %v", startIdx.Add(-time.Hour*9))
		return true
	}
	return false
}

func (s MemberSheetService) ReadSheet(sheet domain.Sheet) (*sheets.ValueRange, error) {
	//spreadsheetId = "1umrFMx3D91eSBF8ytRecK3irLm95npNu8LIrGIKmmOc"
	resp, err := s.sheetReader.ReadSpreadSheet(sheet)
	if err != nil {
		log.Error().Msgf("시트 정보를 가져오는데 실패했습니다. %v", err)
		return nil, err
	}
	return resp, nil
}

func (s MemberSheetService) UpdatePaidMember(resp *sheets.ValueRange) {
	ctx := context.Background()
	//members := make([]*domain.Member, len(resp.Values))
	for i, row := range resp.Values {
		if i == 0 || len(row) < 41 {
			continue
		}
		// 입금 내역 없는 사람도 문자 보냄.
		//if row[39] == nil || row[39] == "" {
		//	continue
		//}
		//if row[40] == nil || row[40] == "" {
		//	continue
		//}
		//members = append(members, &domain.Member{
		//	Id:        i,
		//	PaidAt:    row[35].(string),
		//	PayAmount: row[36].(float64),
		//})
		payAmount, err := strconv.ParseFloat(row[40].(string), 64)
		if err != nil {
			continue
		}
		memberId := eventId + i
		member := &domain.Member{
			Id:        memberId,
			PaidAt:    row[39].(string),
			PayAmount: payAmount,
		}
		num, err := s.task.GetOneByRowNum(ctx, memberId)
		if err != nil {
			log.Error().Msgf("멤버 정보를 가져오는데 실패했습니다. %s %v", row[1], err)
			continue
		}
		// db에 있는 멤버 결제 정보와 비교
		if num.PaidAt == "" || num.PayAmount == 0 {
			err := s.task.UpdateAsPaid(ctx, member)
			if err != nil {
				log.Error().Msgf("멤버 정보를 업데이트하는데 실패했습니다. %v", err)
			}
			log.Info().Msgf("멤버 회비 납부 정보를 업데이트했습니다. 멤버 행 %d", member.Id)
			log.Info().Msgf("입금(날짜/금액): %s/%f", member.PaidAt, member.PayAmount)
			// QR코드 전송을 요청한다.
			go member.ReadyQrTask()
		}
	}
}

func (s MemberSheetService) GetReadyNewRegisters(resp *sheets.ValueRange, startIdx time.Time) chan domain.Member {
	info := make(chan domain.Member)
	go func() {
		// 데이터 출력
		if len(resp.Values) == 0 {
			log.Info().Msg("No data found.")
			return
		} else {
			info = s.defineNewRegister(resp, startIdx, info)
		}
	}()
	return info
}

func (s MemberSheetService) defineNewRegister(resp *sheets.ValueRange, startIdx time.Time, info chan domain.Member) chan domain.Member {
	for i, row := range resp.Values {
		if i == 0 || len(row) < 35 {
			continue
		}
		if convertDate, err := s.convertDate(row[0]); err != nil {
			continue
		} else if s.needToSkip(convertDate, startIdx) {
			continue
		} else {
			member := &domain.Member{
				Id:           eventId + i,
				RegisteredAt: convertDate,
				Name:         row[1].(string),
				Gender:       row[2].(string),
				Generation:   row[3].(string),
				Corps:        row[4].(string),
				Region:       row[5].(string),
				Phone:        row[6].(string),
				Group:        row[7].(string),
				Food:         row[34].(string),
			}
			member.CalcDues()
			member.RecordTask("IDLE")
			info <- *member
		}
	}
	return info
}

func (s MemberSheetService) Handle(member *domain.Member) error {
	s.smsService.SendMessage(&Message{Member: *member})
	member.Record.SetState("SUCCESS")

	return nil
}
