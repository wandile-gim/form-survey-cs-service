package service

import (
	"errors"
	"form-survey-cs-service/internal/domain"
)

type MemberSheetService struct {
	sheetReader *domain.SheetReader
	smsService  Notifier
}

func NewMemberSheetService(s *domain.SheetReader, notifier Notifier) *MemberSheetService {
	return &MemberSheetService{
		s, notifier,
	}
}

func (s *MemberSheetService) ReadSheet(sheet domain.Sheet) (domain.Information, error) {
	//spreadsheetId = "1umrFMx3D91eSBF8ytRecK3irLm95npNu8LIrGIKmmOc"
	resp, err := s.sheetReader.ReadSpreadSheet(sheet)
	if err != nil {
		return domain.Information{}, err
	}

	var group *domain.Information
	// 데이터 출력
	if len(resp.Values) == 0 {
		return domain.Information{}, errors.New("no data found")
	} else {
		// from B, C, D, E, F, G, H
		group = &domain.Information{}
		for i, row := range resp.Values {
			if i == 0 {
				continue
			}
			member := domain.Member{
				Name:       row[0].(string),
				Gender:     row[1].(string),
				Generation: row[2].(string),
				Corps:      row[3].(string),
				Region:     row[4].(string),
				Phone:      row[5].(string),
				Group:      row[6].(string),
			}
			member.DefineTransitCode()
			group.AddMember(member)
		}
	}
	return *group, nil
}

func (s *MemberSheetService) Handle(information domain.Information) error {
	s.smsService.SendMessage(Message{})
	return nil
}
