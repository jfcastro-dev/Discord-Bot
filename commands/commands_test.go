package commands

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatScheduleMessage(t *testing.T) {
	var timestamp = time.Now()
	timeStr := timestamp.Format("Monday 1/2 3:04 PM")
	testCases := []struct {
		name      string
		message   string
		timestamp time.Time
		expected  string
	}{
		{
			name:      "Accepted Test Case",
			message:   "Baldur's Gate 3",
			timestamp: timestamp,
			expected:  fmt.Sprintf("Thumbs up this message if you'd like to partake in Baldur's Gate 3 on %s.", timeStr),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := formatScheduleMessage(tc.message, tc.timestamp)
			if output != tc.expected {
				t.Errorf("ParseSchedule(%q, %q) = %q; want %q", tc.message, tc.timestamp, output, tc.expected)
			}
		})
	}
}

func TestGetHelpMessage(t *testing.T) {
	result, _ := formatHelpMessage()
	if result == "" {
		t.Error("formatHelpMessage() = \"\"; want non-empty string")
	}
}
