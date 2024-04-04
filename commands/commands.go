package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jfcastro-dev/discord-bot/constants"
	"github.com/jfcastro-dev/discord-bot/notifications"
)

type Bot struct {
	Notifications []notifications.Notification
}

func (b *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(m.Content)), constants.BOT_PREFIX) {
		parseMessage(s, m)
	}
}

// parseMessage takes in the content and sends it to the appropriate handler.
func parseMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	command_arr := strings.Fields(m.Content)
	if len(command_arr) > 1 {
		switch command_arr[1] {
		case constants.SCHEDULE:
			timestamp, err := parseTime(strings.Join(command_arr[2:4], " "))
			activity := strings.Join(command_arr[4:], " ")

			if err != nil {
				log.Println(err)
				return
			}
			duration := time.Until(timestamp)
			if duration < 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "Cannot schedule an event in the past.")
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			message := formatScheduleMessage(activity, timestamp)
			sentMessage, err := s.ChannelMessageSend(m.ChannelID, message)

			log.Println("Scheduled message sent ", timestamp)

			if err != nil {
				log.Println(err)
				return
			}
			notification := &notifications.Notification{
				ChannelID: sentMessage.ChannelID,
				Event:     activity,
				Timestamp: timestamp,
				MessageID: sentMessage.ID,
			}
			notifications.SaveNotification(*notification)

			time.AfterFunc(duration, func() {
				notifications.SendNotificationToUsers(s, m.ChannelID, sentMessage.ID, activity)
			})
		default:
			message, _ := formatHelpMessage()
			s.ChannelMessageSend(m.ChannelID, message)
		}
	}
}

// parseTime takes content that matches the time format and returns a time.Time object.
func parseTime(content string) (time.Time, error) {
	layout := "1/2 3:04PM"

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.Parse(layout, content)
	if err != nil {
		return time.Time{}, err
	}

	year := time.Now().In(location).Year()
	t = time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, location)

	return t, nil

}

// formatScheduleMessage returns a string with the schedule message.
func formatScheduleMessage(activity string, time time.Time) string {
	timestamp := time.Format("Monday 1/2 3:04 PM")
	message := fmt.Sprintf(constants.MESSAGE_SCHEDULE, activity, timestamp)
	return message
}

// GetHelpMessage returns a string with the help message as a catch all.
func formatHelpMessage() (string, string) {
	scheduleCommand := fmt.Sprintf("%s %s Jackbox 3/10 7:00PM - Schedules Jackbox at 7:00 PM EST on 3/10", constants.BOT_PREFIX, constants.SCHEDULE)
	helpText := fmt.Sprintf("Here is a list of commands - \n%s", scheduleCommand)
	return helpText, ""
}
