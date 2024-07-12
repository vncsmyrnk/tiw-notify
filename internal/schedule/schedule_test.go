// vim: noexpandtab

package schedule_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/vncsmyrnk/tiwnotify/internal/schedule"
	"github.com/vncsmyrnk/tiwnotify/internal/utils"
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

func TestNewJobByTime_ShouldBeOk(t *testing.T) {
	expected := &schedule.Job{Timer: schedule.NewTimer(time.Second), Task: func() {}}
	job, err := schedule.NewJobByTime(time.Now().Add(time.Second), func() {})
	if err != nil {
		t.Error("error when creating job:", err)
	}
	assert.True(t, cmp.Equal(*expected, *job, utils.IgnoreFuncFields(), cmpopts.IgnoreUnexported(schedule.TimeJobTimer{})))
}
