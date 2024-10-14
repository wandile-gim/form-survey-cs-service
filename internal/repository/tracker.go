package repository

import (
	"context"
	"fmt"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository/ent"
	"form-survey-cs-service/internal/repository/ent/tracker"
	"time"
)

type TrackerRepository struct {
	ent *ent.Client
}

func (t *TrackerRepository) GetLastTask(service string) ([]domain.Tracker, error) {
	var c []domain.Tracker
	if service == "member" {
		all, err := t.ent.Tracker.Query().
			Where(tracker.ServiceEQ(tracker.ServiceMEMBER)).
			Order(ent.Desc(tracker.FieldLastOne)).
			Limit(1).All(context.Background())
		if err != nil {
			return nil, err
		}
		if len(all) > 0 {
			all, _ = t.ent.Tracker.Query().Where(tracker.ServiceEQ(tracker.ServiceMEMBER)).Where(tracker.LastOneGT(all[0].LastOne)).All(context.Background())
			for _, v := range all {
				c = append(c, domain.Tracker{
					LastOne: v.LastOne,
					Service: v.Service.String(),
				})
			}
		}
	}
	return c, nil
}

func (t *TrackerRepository) CreateTracker(service string, lastOne time.Time) error {
	var s tracker.Service
	if service == "member" {
		s = tracker.ServiceMEMBER
	}
	ctx := context.Background()

	// 먼저 존재 여부를 확인하고 트래커를 생성
	if exists, _ := t.ent.Tracker.Query().Where(tracker.ServiceEQ(s)).Exist(ctx); !exists {
		_, err := t.ent.Tracker.Create().
			SetService(s).
			SetLastOne(lastOne).
			SetVersion(time.Now()). // 버전을 현재 시간으로 설정
			Save(ctx)
		if err != nil {
			return fmt.Errorf("tracker 생성 중 에러: %w", err)
		}
	}

	// 트랜잭션으로 낙관적 락 구현
	err := WithTx(context.Background(), t.ent, func(tx *ent.Tx) error {
		// 최신 버전의 Tracker를 가져옴
		latestTracker, err := tx.Tracker.Query().
			Where(tracker.ServiceEQ(s)).
			Order(ent.Desc(tracker.FieldVersion)). // 최신 버전 먼저
			Limit(1).
			First(ctx)
		if err != nil {
			return fmt.Errorf("최신 Tracker 가져오기 실패: %w", err)
		}

		// 최신 버전의 Tracker와 비교하여 동시성 문제 확인
		targetTracker, err := tx.Tracker.Query().
			Where(tracker.ServiceEQ(s)).
			Order(ent.Desc(tracker.FieldVersion)).
			Limit(1).
			First(ctx)
		if err != nil {
			return fmt.Errorf("target Tracker 가져오기 실패: %w", err)
		}

		// 최신 버전과 타겟 버전 비교
		if targetTracker.Version != latestTracker.Version {
			return fmt.Errorf("버전 충돌 발생: 최신 버전이 아님")
		}

		// 최신 버전이 맞다면 업데이트 수행
		_, err = tx.Tracker.UpdateOne(targetTracker).
			SetLastOne(lastOne).    // 필요한 필드 업데이트
			SetVersion(time.Now()). // 버전을 최신 시간으로 갱신
			Save(ctx)
		if err != nil {
			return fmt.Errorf("tracker 업데이트 실패: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("트랜잭션 중 에러: %w", err)
	}

	return nil
}
func NewTrackerRepository(ent *ent.Client) TrackerRepository {
	return TrackerRepository{ent: ent}
}
