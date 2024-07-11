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
	timer       time.Timer
}

func (a Appointment) Notify() {
	notification.Notify("Appointment reminder", a.Description)
}

func (a *Appointment) ScheduleNotification() {
	a.timer = *schedule.AddJob(schedule.Job{Time: a.Time, Task: func() { a.Notify() }})
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
	return Appointment{
		Time:        time,
		Description: parts[1],
	}, nil
}

func ReadAppointmentsFromFile(fileName string) ([]Appointment, error) {
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
