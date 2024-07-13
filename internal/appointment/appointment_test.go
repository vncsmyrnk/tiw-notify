// vim: noexpandtab

package appointment_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vncsmyrnk/tiwnotify/internal/appointment"
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
