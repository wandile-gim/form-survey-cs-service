package main

import (
	"context"
	"fmt"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository"
	"form-survey-cs-service/internal/service"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"html"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	//config.DefaultSetupFromEnv()
	repository.Migration()
	db := repository.Open()

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
		//SpreadsheetId: "1V0l_JA6LQ7EuRt0Zb7LbPpdqzOUdaU4RTPmXnUDVP58",
		SpreadsheetId: "1umrFMx3D91eSBF8ytRecK3irLm95npNu8LIrGIKmmOc",
		Name:          "설문지 응답 시트1",
		Range:         "A:AO",
		Begin:         0,
	}
	sms := service.NewSMSService()
	tDb := repository.NewEntTaskRepository(db)
	tracker := repository.NewTrackerRepository(db)
	taskRepo := repository.NewEntTaskRepository(db)
	s := service.NewMemberSheetService(r, taskRepo, tracker, sms)
	workerService := service.NewWorkerService(tDb, db, s, tracker)

	worker := service.NewWorker()

	go func() {
		for {
			select {
			case l := <-sms.LogChan:
				log.Info().Msg(l)
			}
		}
	}()

	saveTick := time.NewTicker(10 * time.Second)
	taskTick := time.NewTicker(10 * time.Second)
	updateTick := time.NewTicker(10 * time.Second)
	go worker.Run(ctx)

	for {
		select {
		case <-saveTick.C:
			// 시트를 읽어서 태스크를 저장하는 함수
			workerService.SaveTasks(sheet)
		case <-taskTick.C:
			// 작업 기록 테이블에서 상태가 idle인 task를 가져온다.
			tasks := workerService.LoadTasks(sheet)
			go func(tasks chan *service.Task) {
				defer close(tasks)
				for t := range tasks {
					worker.Ready(t)
				}
			}(tasks)
		case <-updateTick.C:
			err := workerService.UpdateTask(sheet)
			if err != nil {
				log.Error().Msgf("UpdateTask: %v", err)
			}
		}
	}
}
