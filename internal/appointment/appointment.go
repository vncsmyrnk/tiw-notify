// vim: noexpandtab

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

func NewAppointment(time time.Time, description string, jobScheduler schedule.JobScheduler) (*Appointment, error) {
	appointment := &Appointment{
		Time:        time,
		Description: description,
	}

	job, err := schedule.NewJobByTime(appointment.Time, func() { appointment.Notify() })
	if err != nil {
		return nil, err
	}

	jobScheduler.AddJob(*job)
	return appointment, nil
}

func (a Appointment) Notify() {
	notifier := notification.BeeepNotifier{}
	notifier.Notify("Appointment reminder", a.Description)
}

func (a Appointment) ScheduleNotificationJob(jobScheduler schedule.JobScheduler) error {
	job, err := schedule.NewJobByTime(a.Time, func() { a.Notify() })
	if err != nil {
		return nil
	}

	jobScheduler.AddJob(*job)
	return nil
}

func ParseAppointmentFromString(str string) (Appointment, error) {
	parts := strings.SplitN(str, " ", 2)
	now := time.Now()
	time, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), parts[0]))
	if err != nil {
		return Appointment{}, err
	}
	return Appointment{Time: time, Description: parts[1]}, nil
}

func ScheduleAppointmentNotificationsFromFile(fileName string) ([]Appointment, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []Appointment{}, err
	}

	var appointments []Appointment
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, err := ParseAppointmentFromString(line)
		if err != nil {
			continue
		}
		appointments = append(appointments, a)
	}
	return appointments, nil

}
