// vim: noexpandtab

//go:generate mockgen -source=schedule.go -destination=mocks/mock_schedule.go -package=mocks

package schedule

import (
	"errors"
	"fmt"
	"time"
)

type JobScheduler interface {
	AddJob(Job)
	StopAllJobs()
}

type Job struct {
	Name  string
	Task  func()
	Timer JobTimer
}

func (j Job) Matches(x any) bool {
	return j.Name == x.(Job).Name
}

func (j Job) String() string {
	return j.Name
}

func NewJobByTime(t time.Time, task func()) (*Job, error) {
	if timeToJob := time.Until(t); timeToJob > 0 {
		timer := NewTimer(timeToJob)
		return &Job{Name: fmt.Sprintf("Job to be done at %v", t), Timer: timer, Task: task}, nil
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
