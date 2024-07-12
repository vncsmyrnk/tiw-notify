// vim: noexpandtab

package schedule

import (
	"errors"
	"time"
)

type Job struct {
	Time time.Time
	Task func()
}

func AddJob(job Job) (*time.Timer, error) {
	if timeToJob := time.Until(job.Time); timeToJob > 0 {
		timer := time.NewTimer(timeToJob)
		go func() {
			<-timer.C
			job.Task()
		}()
		return timer, nil
	} else {
		return nil, errors.New("Given time already passed.")
	}
}
