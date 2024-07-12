// vim: noexpandtab

package schedule_test

import (
	"testing"
	"time"

	"github.com/vncsmyrnk/tiwnotify/internal/schedule"
)

type mockTimer struct {
	CChan chan time.Time
}

func (m *mockTimer) C() <-chan time.Time {
	return m.CChan
}

func (m *mockTimer) Stop() bool {
	return true
}

func NewMockTimer(d time.Duration) schedule.JobTimer {
	return &mockTimer{CChan: make(chan time.Time)}
}

func TestAddJob(t *testing.T) {
	scheduler := &schedule.Schedule{}
	timer := NewMockTimer(1 * time.Second)

	executed := false
	job := schedule.Job{
		Timer: timer,
		Task: func() {
			executed = true
		},
	}

	scheduler.AddJob(job)

	mockTimer := timer.(*mockTimer)
	mockTimer.CChan <- time.Now()

	if !executed {
		t.Errorf("expected task to be executed, but it was not")
	}
}
