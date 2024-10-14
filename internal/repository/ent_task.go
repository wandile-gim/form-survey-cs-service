package repository

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"form-survey-cs-service/internal/repository/ent"
	"form-survey-cs-service/internal/repository/ent/task"
	"form-survey-cs-service/internal/repository/ent/taskrecord"
	"github.com/rs/zerolog/log"
	"time"
)

type EntTaskRepository struct {
	ent *ent.Client
}

func (e EntTaskRepository) GetLastTaskInIdleTimeStamp(ctx context.Context, typeString string) (time.Time, error) {
	var lastOne time.Time
	err := WithTx(ctx, e.ent, func(tx *ent.Tx) error {
		one, err := tx.Task.Query().
			Where(task.TypeEQ(typeString)).
			WithTaskRecord(func(query *ent.TaskRecordQuery) {
				query.Where(taskrecord.StateEQ("IDLE"))
			}).Order(ent.Desc(task.FieldRegisteredAt)).Limit(1).First(ctx)
		if err != nil {
			log.Error().Msgf("GetLastTaskInIdleTimeStamp: %v", err)
			return err
		}
		lastOne = one.RegisteredAt
		return nil
	})
	if err != nil {
		return time.Time{}, err
	}
	return lastOne, nil
}

func (e EntTaskRepository) FindTaskByRegisteredAtAndNotInStatusWorkBegan(ctx context.Context, registeredAt time.Time) ([]*domain.Member, error) {
	var members []*domain.Member
	err := WithTx(ctx, e.ent, func(tx *ent.Tx) error {
		all, err := tx.Task.Query().
			Where(task.RegisteredAtEQ(registeredAt)).
			WithTaskRecord(func(query *ent.TaskRecordQuery) {
				query.Where(taskrecord.StateNotIn("SUCCESS", "FAILED", "RUNNING"))
			}).All(ctx)
		if err != nil {
			return err
		}
		for _, task := range all {
			recordOnly, _ := task.QueryTaskRecord().Only(ctx)
			member := &domain.Member{
				Id:           task.ID,
				Name:         task.Name,
				Phone:        task.Phone,
				Group:        task.Group,
				Corps:        task.Corps,
				Gender:       task.Gender,
				Generation:   task.Generation,
				Region:       task.Region,
				RegisteredAt: task.RegisteredAt,
				Record: &domain.TaskRecord{
					Id:    recordOnly.ID,
					Retry: recordOnly.RetryCount,
				},
			}
			//member.DefineTransitCode()
			member.Record.SetState(recordOnly.State)
			members = append(members, member)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (e EntTaskRepository) CreateMemberTask(ctx context.Context, task domain.Member) (domain.Member, error) {
	var member domain.Member
	err := WithTx(ctx, e.ent, func(tx *ent.Tx) error {
		t, err := tx.Task.Create().
			SetName(task.Name).
			SetType(domain.ServiceMEMBER).
			SetRowNum(task.Id).
			SetCorps(task.Corps).
			SetGroup(task.Group).
			SetPhone(task.Phone).
			SetGeneration(task.Generation).
			SetGender(task.Gender).
			SetRegion(task.Region).
			SetRegisteredAt(task.RegisteredAt).Save(ctx)
		if err != nil {
			log.Error().Msgf("CreateMemberTask Error: %v", err)
			return err
		}
		record_id, err := tx.TaskRecord.Create().
			SetTask(t).
			SetState(task.Record.GetState()).Save(ctx)
		if err != nil {
			return err
		}

		m := &domain.Member{
			Id:           t.ID,
			Name:         t.Name,
			Phone:        t.Phone,
			Group:        t.Group,
			Corps:        t.Corps,
			Gender:       t.Gender,
			Generation:   t.Generation,
			Region:       t.Region,
			RegisteredAt: t.RegisteredAt,
		}
		m.RecordTask(task.Record.GetState())
		m.Record.Id = record_id.ID
		member = *m
		return nil
	})
	if err != nil {
		return domain.Member{}, err
	}
	return member, nil
}

func (e EntTaskRepository) GetIdleStateMemberTasks(ctx context.Context, typeString string) ([]*domain.Member, error) {
	var tasks []*domain.Member
	err := WithTx(context.Background(), e.ent, func(tx *ent.Tx) error {
		taskRecords, err := tx.TaskRecord.
			Query().
			Where(taskrecord.StateEQ("IDLE")).
			Where(taskrecord.HasTask()). // EXISTS: 관계가 있는 TaskRecord만 필터링
			WithTask(func(query *ent.TaskQuery) {
				query.Where(task.TypeEQ("member")) // "tasks"."type" = 'member'
			}).
			All(ctx)
		if err != nil {
			return err
		}
		// Eager Loading 된 Task를 바로 사용 (다시 QueryTask 호출하지 않음)
		for _, taskRecord := range taskRecords {
			task := taskRecord.Edges.Task // 이미 Eager Loaded 된 Task 사용
			if task != nil {
				m := &domain.Member{
					Id:           task.ID,
					Name:         task.Name,
					Phone:        task.Phone,
					Group:        task.Group,
					Corps:        task.Corps,
					Gender:       task.Gender,
					Generation:   task.Generation,
					Region:       task.Region,
					RegisteredAt: task.RegisteredAt,
				}
				m.RecordTask(taskRecord.State)
				m.Record.Id = taskRecord.ID
				tasks = append(tasks, m)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (e EntTaskRepository) UpdateTaskState(task *domain.Member) error {
	if task == nil {
		return nil
	}
	return WithTx(context.Background(), e.ent, func(tx *ent.Tx) error {
		_, err := tx.TaskRecord.
			UpdateOneID(task.Record.Id).
			SetState(task.Record.GetState()).
			SetRetryCount(task.Record.Retry).
			Save(context.Background())
		return err
	})
}

func NewEntTaskRepository(ent *ent.Client) TaskRepository {
	return &EntTaskRepository{
		ent: ent,
	}
}
