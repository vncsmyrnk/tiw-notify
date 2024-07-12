// vim: noexpandtab

//go:generate mockgen -source=timer.go -destination=mocks/mock_timer.go -package=mocks

package schedule

import (
	"time"
)

type JobTimer interface {
	C() <-chan time.Time
	Stop() bool
}

type TimeJobTimer struct {
	timer *time.Timer
}

func (t *TimeJobTimer) C() <-chan time.Time {
	return t.timer.C
}

func (t *TimeJobTimer) Stop() bool {
	return t.timer.Stop()
}

func NewTimer(d time.Duration) *TimeJobTimer {
	return &TimeJobTimer{timer: time.NewTimer(d)}
}
