package domain

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"net/http"
)

type SheetServiceInterface interface {
	ReadSpreadSheet(sheet Sheet)
}

type SheetReader struct {
	service *sheets.Service
}

type Sheet struct {
	SpreadsheetId string
	Name          string
	Range         string
	Begin         int
}

func NewSheet(spreadsheetId, name, _range string) *Sheet {
	return &Sheet{
		SpreadsheetId: spreadsheetId,
		Name:          name,
		Range:         _range,
	}
}

func (s *Sheet) GetSheetName() string {
	return fmt.Sprintf("%s!%s", s.Name, s.Range)
}

func NewSheetReader(ctx context.Context, client *http.Client) *SheetReader {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
	return &SheetReader{
		service: srv,
	}
}

func (s *SheetReader) ReadSpreadSheet(sheet Sheet) (*sheets.ValueRange, error) {
	resp, err := s.service.Spreadsheets.Values.Get(sheet.SpreadsheetId, sheet.GetSheetName()).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
