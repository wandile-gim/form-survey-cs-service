package domain

import (
	"context"
	"form-survey-cs-service/internal/config"
	"form-survey-cs-service/internal/repository"
	"testing"
)

func TestMember_ReadyQrTask(t *testing.T) {
	config.DefaultSetupFromEnv()
	repo := repository.NewEntTaskRepository(repository.Open())
	member, err := repo.GetOneByRowNum(context.Background(), 1)
	if err != nil {
		return
	}

	member.ReadyQrTask()
}
