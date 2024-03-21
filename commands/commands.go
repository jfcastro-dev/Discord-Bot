package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jfcastro-dev/discord-bot/constants"
	"strings"
	"time"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(m.Content)), constants.BOT_PREFIX) {
		message := ParseMessage(m.Content)
		fmt.Println(message)
		s.ChannelMessageSend(m.ChannelID, message)
	}
}

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

// ParseSchedule takes content that matches the Schedule format and returns a message on match.
func ParseSchedule(content string) string {
	schedule_arr := strings.Fields(content)
	layout := "1/2 3:04PM"
	t, err := time.Parse(layout, strings.Join(schedule_arr[1:], " "))
	year := time.Now().Year()

	t = time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())

	if err != nil {
		fmt.Println(err)
		return GetHelpMessage()
	}
	return printSchedule(schedule_arr[0], t)
}

func printSchedule(activity string, time time.Time) string {
	s := time.Format("Monday 1/2 3:04 PM")
	message := fmt.Sprintf("React to this message if you'd like to partake in %s on %s.", activity, s)
	return message
}

// GetHelpMessage returns a string with the help message as a catch all.
func GetHelpMessage() string {
	scheduleCommand := fmt.Sprintf("%s %s Jackbox 3/10 7:00PM - Schedules Jackbox at 7:00 PM EST on 3/10", constants.BOT_PREFIX, constants.SCHEDULE)
	helpText := fmt.Sprintf("Here is a list of commands - \n%s", scheduleCommand)
	return helpText
}
