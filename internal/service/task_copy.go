package service

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"github.com/rs/zerolog/log"
)

// SaveTasks 시트를 읽어서 태스크를 저장하는 함수.
func (t WorkerService) SaveTasks(sheetDomain domain.Sheet) {
	// 인덱스 번호를 줘야해
	startIdx, err := t.repo.GetLastTaskInIdleTimeStamp(context.Background(), domain.ServiceMEMBER)
	if err != nil {
		log.Error().Msgf("태스크 정보를 가져오는데 실패했습니다. %v", err)
	}
	sheet, err := t.sheetService.ReadSheet(sheetDomain, startIdx)
	if err != nil {
		log.Error().Msgf("시트 정보를 가져오는데 실패했습니다. %v", err)
	}

	go func() {
		for {
			select {
			case member, ok := <-sheet:
				if !ok {
					log.Info().Msg("시트 정보를 모두 읽었습니다.")
					return
				}
				_, err := t.repo.CreateMemberTask(context.Background(), member)
				if err != nil {
					//log.Error().Msgf("태스크 정보를 생성하는데 실패했습니다. member:%v %v", member, err)
					continue
				}
			}
		}
	}()
}

func (t WorkerService) LoadTasks(sheetDomain domain.Sheet) chan *Task {
	result := make(chan *Task)

	go func() {
		trackers, err := t.tracker.GetLastTask(domain.ServiceMEMBER)
		if err != nil {
			log.Error().Msgf("트래커 정보를 가져오는데 실패했습니다. %v", err)
			return
		}

		if len(trackers) > 1 {
			// 마지막으로 실행된 정보가 있다면
			// 해당 정보를 기준으로 실행되지 않은 정보만 로드하기.
			// 마지막으로 실행된 정보가 없다면
			// 모든 정보 로드하기.
			// not in running, failed, success
			for _, tracker := range trackers {
				begans, err := t.repo.FindTaskByRegisteredAtAndNotInStatusWorkBegan(context.Background(), tracker.LastOne)
				if err != nil {
					log.Error().Msgf("태스크 정보를 가져오는데 실패했습니다. %v", err)
				}
				for _, began := range begans {
					taskCopy := NewTask(sheetDomain, t.sheetService, t) // 깊은 복사
					taskCopy.member = began
					result <- taskCopy
				}
			}
		}

		loadedTasks, err := t.repo.GetIdleStateMemberTasks(context.Background(), domain.ServiceMEMBER)
		if err != nil {
			log.Error().Msgf("태스크 정보를 가져오는데 실패했습니다. %v", err)
		}
		for _, loadedTask := range loadedTasks {
			taskCopy := NewTask(sheetDomain, t.sheetService, t) // 깊은 복사
			taskCopy.member = loadedTask
			result <- taskCopy
		}
	}()
	return result
}
