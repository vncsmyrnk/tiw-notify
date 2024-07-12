// vim: noexpandtab

//go:generate mockgen -source=schedule.go -destination=mocks/mock_schedule.go -package=mocks

package schedule

import (
	"errors"
	"time"
)

type JobScheduler interface {
	AddJob(Job)
	StopAllJobs()
}

type Job struct {
	Task func()
	Timer JobTimer
}

func NewJobByTime(t time.Time, task func()) (*Job, error) {
	if timeToJob := time.Until(t); timeToJob > 0 {
		timer := NewTimer(timeToJob)
		return &Job{Timer: timer, Task: task}, nil
	} else {
		return nil, errors.New("Given time already passed.")
	}
}

type Schedule struct {
	Jobs []Job
}

func (s *Schedule) AddJob(job Job) {
	s.Jobs = append(s.Jobs, job)
	go func() {
		<-job.Timer.C()
		job.Task()
	}()
}

func (s Schedule) StopAllJobs() {
	for _, j := range s.Jobs {
		j.Timer.Stop()
	}
}
