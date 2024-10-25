package service

import (
	"form-survey-cs-service/internal/domain"
	"google.golang.org/api/sheets/v4"
	"time"
)

type SheetService interface {
	ReadSheet(sheet domain.Sheet) (*sheets.ValueRange, error)
	Handle(information *domain.Member) error
	GetReadyNewRegisters(resp *sheets.ValueRange, startIdx time.Time) chan domain.Member
	UpdatePaidMember(resp *sheets.ValueRange)
}
