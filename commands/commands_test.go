package commands

import (
	"fmt"
	"github.com/jfcastro-dev/discord-bot/constants"
	"testing"
)

func TestParseSchedule(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Accepted Test Case",
			input:    "3/10 7:00PM",
			expected: "React to this message if you'd like to partake in this session.",
		},
		{
			name:     "Rejected Test Case",
			input:    "input2",
			expected: GetHelpMessage(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ParseSchedule(tc.input)
			if output != tc.expected {
				t.Errorf("ParseSchedule(%q) = %q; want %q", tc.input, output, tc.expected)
			}
		})
	}
}

func TestGetHelpMessage(t *testing.T) {
	result := GetHelpMessage()
	if result == "" {
		t.Error("GetHelpMessage() = \"\"; want non-empty string")
	}
}

func TestParseCommand(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Parse Date Command",
			input:    fmt.Sprintf("%s %s 3/10 7:00PM", constants.BOT_PREFIX, constants.SCHEDULE),
			expected: "React to this message if you'd like to partake in this session.",
		},
		{
			name:     "Get Help",
			input:    fmt.Sprintf("%s help", constants.BOT_PREFIX),
			expected: GetHelpMessage(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ParseMessage(tc.input)
			if output != tc.expected {
				t.Errorf("ParseSchedule(%q) = %q; want %q", tc.input, output, tc.expected)
			}
		})
	}
}
