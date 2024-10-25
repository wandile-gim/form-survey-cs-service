package domain

import (
	"form-survey-cs-service/internal/config"
	"testing"
	"time"
)

func TestMember_ReadyQrTask(t *testing.T) {
	config.DefaultSetupFromEnv()

	type fields struct {
		Id           int
		Name         string
		Phone        string
		Group        string
		Corps        string
		Gender       string
		Generation   string
		Region       string
		RegisteredAt time.Time
		PayAmount    float64
		PaidAt       string
		Food         string
		Record       *TaskRecord
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test ReadyQrTask",
			fields: fields{
				Id:           1,
				Name:         "test",
				Phone:        "010-1234-5678",
				Group:        "test",
				Corps:        "test",
				Gender:       "test",
				Generation:   "15",
				Region:       "test",
				RegisteredAt: time.Now(),
				PayAmount:    5000,
				Food:         "test",
				Record:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Member{
				Id:           tt.fields.Id,
				Name:         tt.fields.Name,
				Phone:        tt.fields.Phone,
				Group:        tt.fields.Group,
				Corps:        tt.fields.Corps,
				Gender:       tt.fields.Gender,
				Generation:   tt.fields.Generation,
				Region:       tt.fields.Region,
				RegisteredAt: tt.fields.RegisteredAt,
				PayAmount:    tt.fields.PayAmount,
				PaidAt:       tt.fields.PaidAt,
				Food:         tt.fields.Food,
				Record:       tt.fields.Record,
			}
			i.ReadyQrTask()
		})
	}
}
