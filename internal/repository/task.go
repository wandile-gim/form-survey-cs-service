package repository

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"time"
)

type TaskRepository interface {
	CreateMemberTask(ctx context.Context, task domain.Member) (domain.Member, error)
	GetIdleStateMemberTasks(ctx context.Context, typeString string) ([]*domain.Member, error)
	GetOneByRowNum(ctx context.Context, rowNum int) (*domain.Member, error)

	UpdateTaskState(task *domain.Member) error
	UpdateAsPaid(ctx context.Context, ta *domain.Member) error

	GetLastTaskInIdleTimeStamp(ctx context.Context, typeString string) (time.Time, error)
	FindTaskByRegisteredAtAndNotInStatusWorkBegan(ctx context.Context, registeredAt time.Time) ([]*domain.Member, error)
}
