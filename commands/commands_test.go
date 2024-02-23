package commands

import (
	"testing"
)

func TestParseSchedule(t *testing.T) {
	// Test with a valid date string
	result := ParseSchedule("3/10 7:00PM")
	expected := "Would you like to partake in the session?"
	if result != expected {
		t.Errorf("ParseSchedule(\"3/10 7:00PM\") = %s; want %s", result, expected)
	}

	// Test with an invalid date string
	result = ParseSchedule("invalid date string")
	expected = GetHelpMessage()
	if result != expected {
		t.Errorf("ParseSchedule(\"invalid date string\") = %s; want %s", result, expected)
	}
}

func TestGetHelpMessage(t *testing.T) {
	result := GetHelpMessage()
	if result == "" {
		t.Error("GetHelpMessage() = \"\"; want non-empty string")
	}
}
