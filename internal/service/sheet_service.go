package service

import "form-survey-cs-service/internal/domain"

type SheetService interface {
	ReadSheet(sheet domain.Sheet) (domain.Information, error)
	Handle(information domain.Information) error
}
