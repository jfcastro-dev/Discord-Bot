package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jfcastro-dev/discord-bot/commands"
	"log"
	"os"
	"os/signal"
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
	session.AddHandler(commands.MessageHandler)
	err = session.Open()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
}
