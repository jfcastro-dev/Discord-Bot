package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/jfcastro-dev/discord-bot/commands"
	"github.com/jfcastro-dev/discord-bot/notifications"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	notificationArr, err := notifications.LoadNotifications()
	if err != nil {
		log.Fatal(err)
	}

	bot := &commands.Bot{
		Notifications: notificationArr,
	}

	/*for _, notification := range notificationObject.Notifications {
		if notification.Event != "" {
			time.AfterFunc(notification.Timestamp.Sub(time.Now()), func() {
				notifications.SendNotificationToUsers(session, &notification)
			})
		}
	}*/

	session.AddHandler(bot.MessageHandler)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	log.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}
