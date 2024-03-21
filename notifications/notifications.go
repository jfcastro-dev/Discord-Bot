package notifications

import (
	"encoding/json"
	"fmt"
	"github.com/jfcastro-dev/discord-bot/constants"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Notification struct {
	Users     []string  `json:"users"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
	ChannelID string    `json:"channel_id"`
}

type Notifications struct {
	Notifications []Notification `json:"notifications"`
}

/**
 * LoadNotifications loads the notifications from the database.
 */
func LoadNotifications() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting current file path")
		return
	}

	appDir := filepath.Dir(filepath.Dir(filename))

	path := filepath.Join(appDir, constants.DATA_PATH, constants.NOTIFICATIONS_FILE)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var notifications Notifications

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&notifications)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	registerNotifications(notifications)
}

func registerNotifications(notifications Notifications) {
	for _, notification := range notifications.Notifications {
		fmt.Printf("Register notification %v", notification.Event)
		executeNotifications(notification)
	}
}

func executeNotifications(notification Notification) {
	fmt.Sprintf("Execute notification %s", notification.Event)
}
