package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token = ""
)

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	defer discord.Close()

	return
}

func onMessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author
	if user.ID == session.State.User.ID || user.Bot {
		return
	}

	content := msg.Content

	switch content {
	case "ping", "Ping":
		session.ChannelMessageSend(msg.ChannelID, "Pong!")

	case "pong", "Pong!":
		session.ChannelMessageSend(msg.ChannelID, "Ping!")
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}
