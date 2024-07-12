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
	timer       *time.Timer
}

func New(time time.Time, description string) (*Appointment, error) {
	appointment := &Appointment{
		Time:        time,
		Description: description,
	}
	timer, err := schedule.AddJob(schedule.Job{Time: appointment.Time, Task: func() { appointment.Notify() }})
	if err != nil {
		return nil, err
	}
	appointment.timer = timer
	return appointment, nil
}

func (a Appointment) Notify() {
	notification.Notify("Appointment reminder", a.Description)
}

func (a Appointment) StopJob() {
	a.timer.Stop()
}

func ParseAppointmentFromString(str string) (Appointment, error) {
	parts := strings.SplitN(str, " ", 2)
	now := time.Now()
	time, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), parts[0]))
	if err != nil {
		return Appointment{}, err
	}
	appointment, err := New(time, parts[1])
	if err != nil {
		return Appointment{}, err
	}
	return *appointment, nil
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
