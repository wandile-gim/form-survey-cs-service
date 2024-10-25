package service

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository"
	"form-survey-cs-service/internal/repository/ent"
	"github.com/rs/zerolog/log"
)

type WorkerUseCase interface {
	SaveTasks(sheetDomain domain.Sheet)
	LoadTasks(sheetDomain domain.Sheet) chan *Task
	UpdateTask(sheet domain.Sheet) error
	TaskFailed(ctx context.Context, member *domain.Member, task *Task) string
	TaskSuccess(information *domain.Member) string
}

type WorkerService struct {
	tx           *ent.Client
	sheetService SheetService
	repo         repository.TaskRepository
	tracker      repository.TrackerRepository
}

//func (t WorkerService) LoadTasks(sheetDomain domain.Sheet) chan *Task {
//	var startIdx time.Time
//	result := make(chan *Task)
//
//	go func() {
//		task := NewTask(sheetDomain, t.sheetService, t)
//		trackers, err := t.tracker.GetLastTask(domain.ServiceMEMBER)
//		if err != nil {
//			log.Error().Msgf("트래커 정보를 가져오는데 실패했습니다. %v", err)
//			return
//		}
//
//		if len(trackers) > 1 {
//			// 마지막으로 실행된 정보가 있다면
//			// 해당 정보를 기준으로 실행되지 않은 정보만 로드하기.
//			// 마지막으로 실행된 정보가 없다면
//			// 모든 정보 로드하기.
//			// not in running, failed, success
//			for _, tracker := range trackers {
//				begans, err := t.repo.FindTaskByRegisteredAtAndNotInStatusWorkBegan(context.Background(), tracker.LastOne)
//				if err != nil {
//					log.Error().Msgf("태스크 정보를 가져오는데 실패했습니다. %v", err)
//				}
//				for _, began := range begans {
//					taskCopy := NewTask(sheetDomain, t.sheetService, t) // 깊은 복사
//					taskCopy.member = began
//					result <- taskCopy
//				}
//			}
//		} else if len(trackers) == 1 {
//			startIdx = trackers[0].LastOne
//		}
//		// 인덱스 번호를 줘야해
//		t.repo.FindTaskByRegisteredAtAndNotInStatusWorkBegan()
//		sheet, err := task.service.ReadSheet(sheetDomain, startIdx)
//		if err != nil {
//			log.Error().Msgf("시트 정보를 가져오는데 실패했습니다. %v", err)
//			return
//		}
//		go func() {
//			for {
//				select {
//				case member, ok := <-sheet:
//					if !ok {
//						log.Info().Msg("시트 정보를 모두 읽었습니다.")
//						return
//					}
//					_, err := t.repo.CreateMemberTask(context.Background(), member)
//					if err != nil {
//						//log.Error().Msgf("태스크 정보를 생성하는데 실패했습니다. member:%v %v", member, err)
//						continue
//					}
//				}
//			}
//		}()
//	}()
//	return result
//}

func (t WorkerService) TaskFailed(ctx context.Context, info *domain.Member, task *Task) string {
	// 태스크 레코드에 실패 정보 저장.
	if info == nil {
		return ""
	}
	info.Record.IncreaseRetry()
	task.failed = info.Record.Retry

	err := repository.WithTx(ctx, t.tx, func(tx *ent.Tx) error {
		err := t.repo.UpdateTaskState(info)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return info.Record.GetState()
	}
	return info.Record.GetState()
}

func (t WorkerService) TaskSuccess(information *domain.Member) string {
	// 태스크 레코드에 성공 정보 저장.
	if information == nil {
		return ""
	} else {
		err := t.repo.UpdateTaskState(information)
		if err != nil {
			log.Error().Msgf("태스크 정보를 업데이트하는데 실패했습니다. %v", err)
			return ""
		} else {
			return information.Record.GetState()
		}
	}
}

func NewWorkerService(
	taskRepository repository.TaskRepository,
	tx *ent.Client,
	sheetService SheetService,
	trackerRepository repository.TrackerRepository,
) WorkerUseCase {
	return &WorkerService{
		tx:           tx,
		sheetService: sheetService,
		repo:         taskRepository,
		tracker:      trackerRepository,
	}
}
