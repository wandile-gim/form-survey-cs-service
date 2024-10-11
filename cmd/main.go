package main

import (
	"context"
	"fmt"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/service"
	"html"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		// 토큰 저장 비즈니스 로직 작성
	})
	go http.ListenAndServe(":8080", nil)

	// OAuth2 클라이언트 생성 (토큰이 유효하지 않으면 자동으로 갱신)
	client := domain.OAuthForSheet()

	// Google Sheets API 클라이언트 생성
	r := domain.NewSheetReader(ctx, client)
	sheet := domain.Sheet{
		SpreadsheetId: "1umrFMx3D91eSBF8ytRecK3irLm95npNu8LIrGIKmmOc",
		Name:          "설문지 응답 시트1",
		Range:         "B:H",
	}
	s := service.NewMemberSheetService(r)
	info, _ := s.ReadSheet(sheet)
	log.Println(info)
}
