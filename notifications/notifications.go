package notifications

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jfcastro-dev/discord-bot/constants"
)

type Notification struct {
	ChannelID string    `json:"channel_id"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
	MessageID string    `json:"message_id"`
}

type Wrapper struct {
	Notifications []Notification `json:"notifications"`
}

// Wrapper function for saveNotification that builds the Notification array to be saved.
func SaveNotification(notification Notification) error {
	notifications, err := LoadNotifications()
	if err != nil {
		return err
	}
	notifications = append(notifications, notification)
	return saveNotifications(notifications)
}

// SaveNotifications saves the notifications to the data source.
func saveNotifications(notifications []Notification) error {
	// Convert the notifications to JSON
	data, err := json.Marshal(notifications)
	if err != nil {
		return err
	}

	// Create the full file path
	fullPath := filepath.Join(constants.DATA_PATH, constants.NOTIFICATIONS_FILE)

	// Write the data to the file
	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// DeleteNotification deletes a notification from the list of notifications. Used after the event has started.
func DeleteNotification(messageID string) {
	notifications, err := LoadNotifications()
	if err != nil {
		log.Println("Error loading notifications:", err)
		return
	}
	for i, notification := range notifications {
		if notification.MessageID == messageID {
			// Remove the notification from the array
			notifications = append(notifications[:i], notifications[i+1:]...)
			break
		}
	}
	saveNotifications(notifications)
}

// LoadNotifications loads the notifications from the data source.
func LoadNotifications() ([]Notification, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("error getting current file path")
	}

	appDir := filepath.Dir(filepath.Dir(filename))

	path := filepath.Join(appDir, constants.DATA_PATH, constants.NOTIFICATIONS_FILE)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notifications []Notification

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&notifications)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

// Helper function used to check if an item is in a slice.
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// getUserReacts gets the users who reacted to a message with a thumbs up.
func getUserReacts(s *discordgo.Session, channelID string, messageID string) ([]string, error) {
	reactions, err := s.MessageReactions(channelID, messageID, constants.REACT_EMOJI, constants.REACT_LIMIT, "", "")
	if err != nil {
		return nil, err
	}

	var users []string
	for _, reaction := range reactions {
		userID := reaction.ID
		if !contains(users, userID) {
			users = append(users, userID)
		}
	}
	return users, nil
}

// SendNotificationToUsers sends a private message to all users who reacted to a message.
func SendNotificationToUsers(s *discordgo.Session, channelID string, messageID string, event string) {
	message := fmt.Sprintf("Event: %s is starting now. Have fun!", event)
	log.Println("Event Starting", event)
	users, err := getUserReacts(s, channelID, messageID)
	if err != nil {
		log.Println("Error getting reactions:", err)
		return
	}

	for _, userID := range users {
		channel, err := s.UserChannelCreate(userID)
		log.Println("Sending message to user:", userID)
		if err != nil {
			log.Println("Error creating channel with user:", err)
			continue
		}

		_, err = s.ChannelMessageSend(channel.ID, message)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}
