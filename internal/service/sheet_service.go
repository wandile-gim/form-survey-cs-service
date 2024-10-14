package service

import (
	"form-survey-cs-service/internal/domain"
	"time"
)

type SheetService interface {
	ReadSheet(sheet domain.Sheet, startIdx time.Time) (<-chan domain.Member, error)
	Handle(information *domain.Member) error
}
