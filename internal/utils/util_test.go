// vim: noexpandtab

package utils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vncsmyrnk/tiwnotify/internal/utils"
)

func TestHourMinuteStringToTime_ShouldBeOk(t *testing.T) {
	testCases := []struct {
		name string
		input string
	}{
		{"Time1", "09:05"},
		{"Time2", "23:59"},
		{"Time3", "00:00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.HourMinuteStringToTime(tc.input)
			if err != nil {
			  t.Error("error parsing hour to time.Time", err)
			}

			now := time.Now()
			expected, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), tc.input))
			if err != nil {
			  t.Error("error creating expected value for test", err)
			}

			assert.Equal(t, expected, result)
		})
	}
}

func TestHourMinuteStringToTime_ShouldErr(t *testing.T) {
	testCases := []struct {
		name string
		input string
	}{
		{"Time1", "0905"},
		{"Time2", "12:02pm"},
		{"Time3", "two o'clock"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := utils.HourMinuteStringToTime(tc.input)
			if err == nil {
			  t.Error("an error should be return and it was not", err)
			}
		})
	}
}
