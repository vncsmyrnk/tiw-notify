// vim: noexpandtab

package schedule

import (
	"time"
)

type Job struct {
	Time time.Time
	Task func()
}

func AddJob(job Job) *time.Timer {
	timer := time.NewTimer(time.Until(job.Time))
	go func() {
		<-timer.C
		job.Task()
	}()
	return timer
}
