// vim: noexpandtab

//go:generate mockgen -source=appointment.go -destination=mocks/mock_appointment.go -package=mocks

package appointment

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vncsmyrnk/tiwnotify/internal/notification"
	"github.com/vncsmyrnk/tiwnotify/internal/schedule"
)

type Appointment struct {
	Time        time.Time
	Description string
}

func NewAppointmentFromString(str string) (Appointment, error) {
	parts := strings.SplitN(str, " ", 2)
	now := time.Now()
	time, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), parts[0]))
	if err != nil {
		return Appointment{}, err
	}
	return Appointment{Time: time, Description: parts[1]}, nil
}

type AppointmentScheduler interface {
	ScheduleFromFile(string) error
}

type AppointmentSchedule struct {
	Scheduler schedule.JobScheduler
	Notifier notification.Notifier
}

func (as AppointmentSchedule) ScheduleFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, err := NewAppointmentFromString(line)
		if err != nil {
			continue
		}
		job, err := schedule.NewJobByTime(a.Time, func() { as.Notifier.Notify("Appointment reminder", a.Description) })
		if err != nil {
			continue
		}
		as.Scheduler.AddJob(*job)
	}
	return nil
}
