package domain

import "sync"

type TaskRecord struct {
	*sync.Mutex
	Id    int
	state string
	Retry int
}

func (t *TaskRecord) SetState(state string) {
	t.state = state
}

func (t *TaskRecord) IncreaseRetry() {
	defer t.Unlock()
	t.Lock()
	t.Retry += 1
	if t.Retry > 3 {
		t.SetState("failed")
	}
}

func (t *TaskRecord) GetState() string {
	return t.state
}
