package service

import (
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

type MemberSheetService struct {
	sheetReader *domain.SheetReader
	tracker     repository.TrackerRepository
	smsService  Notifier
}

func NewMemberSheetService(s *domain.SheetReader, tracker repository.TrackerRepository, notifier Notifier) MemberSheetService {
	return MemberSheetService{
		s, tracker, notifier,
	}
}

func (s MemberSheetService) convertDate(row []interface{}) (time.Time, error) {
	var registeredAt time.Time
	if strings.Contains(row[0].(string), "오후") {
		row[0] = strings.Replace(row[0].(string), "오후", "PM", 1)
	} else if strings.Contains(row[0].(string), "오전") {
		row[0] = strings.Replace(row[0].(string), "오전", "AM", 1)
	}
	if registeredAt2, err := time.Parse("2006. 01. 02 PM 3:04:05", row[0].(string)); err == nil {
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
		log.Info().Msgf("최신 스킵 데이터: %v", startIdx.Add(-time.Hour*9))
		return true
	}
	return false
}

func (s MemberSheetService) ReadSheet(sheet domain.Sheet, startIdx time.Time) (<-chan domain.Member, error) {
	//spreadsheetId = "1umrFMx3D91eSBF8ytRecK3irLm95npNu8LIrGIKmmOc"
	info := make(chan domain.Member)
	resp, err := s.sheetReader.ReadSpreadSheet(sheet)
	if err != nil {
		return nil, err
	}
	go func() {
		// 데이터 출력
		if len(resp.Values) == 0 {
			log.Info().Msg("No data found.")
			return
		} else {
			for i, row := range resp.Values {
				if i == 0 || len(row) < 35 {
					continue
				}
				if convertDate, err := s.convertDate(row); err != nil {
					continue
				} else if s.needToSkip(convertDate, startIdx) {
					continue
				} else {
					member := &domain.Member{
						Id:           i,
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
					//member.DefineTransitCode()
					member.RecordTask("IDLE")
					info <- *member
				}
			}
		}
	}()
	return info, nil
}

func (s MemberSheetService) Handle(member *domain.Member) error {
	s.smsService.SendMessage(&Message{Member: *member})
	member.Record.SetState("SUCCESS")

	return nil
}
