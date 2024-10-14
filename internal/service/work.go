package service

import (
	"context"
	"form-survey-cs-service/internal/domain"
	"log"
	"sync"
)

type Task struct {
	failed      int
	result      string
	sheet       *domain.Sheet
	member      *domain.Member
	service     SheetService
	taskService WorkerService
}

func NewTask(sheet domain.Sheet, service SheetService, taskService WorkerService) *Task {
	return &Task{
		sheet:       &sheet,
		service:     service,
		taskService: taskService,
	}
}

type Worker struct {
	sync.Mutex
	ready  chan *Task
	failed chan *Task
}

func NewWorker() *Worker {
	return &Worker{
		ready:  make(chan *Task),
		failed: make(chan *Task),
	}
}

func (w *Worker) Ready(t *Task) {
	w.ready <- t
}

func (w *Worker) Failed(t *Task) {
	if t.failed > 3 {
		log.Printf("task failed: %s", t.result)
		return
	}
	w.failed <- t
}

func (w *Worker) Run(ctx context.Context) {
	for {
		select {
		case t := <-w.ready:
			//log.Printf("task ready: %v", t.member)
			go func(member *domain.Member) {
				err := t.service.Handle(member)
				if err != nil {
					t.taskService.TaskFailed(ctx, t.member, t)
					w.Failed(t)
				} else {
					t.taskService.TaskSuccess(t.member)
					t.taskService.tracker.CreateTracker("member", t.member.RegisteredAt)
				}
			}(t.member)
		case t := <-w.failed:
			log.Printf("실패한 task 재시작: %s", t.sheet.Name)
		}
	}
}
