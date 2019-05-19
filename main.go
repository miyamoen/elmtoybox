package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	config, err := initialize()
	if err != nil {
		log.Fatal("error initializing bot app,", err)
	}

	discord, err := discordgo.New("Bot " + config.token)
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

// Config is 設定値
type Config struct {
	token string
}

func initialize() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	return &Config{token: os.Getenv("BOT_TOKEN")}, nil

}

func onMessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author
	if user.ID == session.State.User.ID || user.Bot {
		return
	}

	content := msg.Content

	switch strings.ToLower(content) {
	case "ping", "ping!":
		session.ChannelMessageSend(msg.ChannelID, "Pong!")

	case "pong", "pong!":
		session.ChannelMessageSend(msg.ChannelID, "Ping!")
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}
