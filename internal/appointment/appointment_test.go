// vim: noexpandtab

package appointment_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vncsmyrnk/tiwnotify/internal/appointment"
	notificationmocks "github.com/vncsmyrnk/tiwnotify/internal/notification/mocks"
	// "github.com/vncsmyrnk/tiwnotify/internal/schedule"
	schedulemocks "github.com/vncsmyrnk/tiwnotify/internal/schedule/mocks"
	// "github.com/vncsmyrnk/tiwnotify/internal/utils"
)

func TestNewAppointmentFromString_ShouldBeOk(t *testing.T) {
	a, err := appointment.NewAppointmentFromString("20:56 Wash the dishes")
	if err != nil {
		t.Error("error when creating appointment:", err)
	}
	now := time.Now()
	expectedTime, _ := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), "20:56"))
	expected := appointment.Appointment{Time: expectedTime, Description: "Wash the dishes"}
	assert.Equal(t, expected, a)
}

func TestNewAppointmentFromString_ShouldFail(t *testing.T) {
	testCases := []struct {
		name string
		input string
	}{
		{"Letter in time", "18:g8 Malformed appointment"},
		{"No time", "Malformed appointment"},
		{"Not enough numbers in time", "22:3 Malformed appointment"},
		{"No colon", "1134 Malformed appointment"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := appointment.NewAppointmentFromString(tc.input)
			assert.True(t, err != nil)
		})
	}
}

func TestScheduleFromFile_ShouldBeOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	timeStr := time.Now().Add(time.Minute).Format("15:04")
	fmt.Print("time used:", timeStr)

	f, err := os.CreateTemp("", "sample")
	if err != nil {
		t.Error("failed to create temporarily file:", err)
		return
	}

	f.WriteString(fmt.Sprintf("%v A random appointment\n", timeStr))
	defer os.Remove(f.Name())

	mockJobScheduler := schedulemocks.NewMockJobScheduler(ctrl)
	mockNotifier := notificationmocks.NewMockNotifier(ctrl)

	// expectedTime, err := utils.HourMinuteStringToTime(timeStr)
	// if err != nil {
	// 	t.Error("failed to create expected time:", err)
	// 	return
	// }

	// expectedJob, err := schedule.NewJobByTime(expectedTime, func() {})
	// if err != nil {
	// 	t.Error("failed to create expected job:", err)
	// 	return
	// }

	mockJobScheduler.EXPECT().AddJob(gomock.Any())
	as := appointment.AppointmentSchedule{Scheduler: mockJobScheduler, Notifier: mockNotifier}
	err = as.ScheduleFromFile(f.Name())
	if err != nil {
		t.Error("error while scheduling appointments", err)
	}
}
