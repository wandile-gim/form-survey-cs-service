package service

import (
	"log"
	"sync"
)

type Task struct {
	failed  int
	result  string
	service SheetService
}

type Worker struct {
	sync.Mutex
	ready  chan *Task
	failed chan *Task
}

func NewTask() *Worker {
	return &Worker{
		ready:  make(chan *Task),
		failed: make(chan *Task),
	}
}

func (w *Worker) Ready(t *Task) {
	w.ready <- t
}

func (w *Worker) Failed(t *Task) {
	defer w.Unlock()
	w.Lock()
	t.failed++
	if t.failed > 3 {
		log.Printf("task failed: %s", t.result)
	}
	w.failed <- t
}

func (w *Worker) Run() {
	for {
		select {
		case t := <-w.ready:
			log.Printf("task ready: %s", t.result)
			// read sheets and set information
			// define task
			info, err := t.service.ReadSheet()
			// distribute task
			t.service.Handle(info)
		}
	}
}
