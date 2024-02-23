package commands

import (
	"fmt"
	"github.com/jfcastro-dev/discord-bot/constants"
	"strings"
	"time"
)

// ParseMessage takes in the content and sends it to the appropriate handler.
func ParseMessage(content string) string {
	command_arr := strings.Fields(content)
	if command_arr != nil && len(command_arr) > 1 {
		if command_arr[0] == constants.BOT_PREFIX {
			switch command_arr[1] {
			case constants.SCHEDULE:
				return ParseSchedule(strings.Join(command_arr[2:], " "))
			}
		}
	}
	return GetHelpMessage()
}

// ParseScheduler takes content that matches the Schedule format and returns a message on match.
func ParseSchedule(content string) string {
	layout := "1/1 2:00PM"
	_, err := time.Parse(layout, content)
	if err != nil {
		fmt.Println(err)
		return GetHelpMessage()
	}
	return "React to this message if you'd like to partake in this session."
}

// GetHelpMessage returns a string with the help message as a catch all.
func GetHelpMessage() string {
	scheduleCommand := fmt.Sprintf("%s %s Jackbox 3/10 7:00PM - Schedules Jackbox at 7:00 PM EST on 3/10", constants.BOT_PREFIX, constants.SCHEDULE)
	helpText := fmt.Sprintf("Here is a list of commands - \n%s", scheduleCommand)
	return helpText
}
