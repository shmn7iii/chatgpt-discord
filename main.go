package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/shmn7iii/chatgpt-discord/internal/chatgpt"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	s.AddHandler(chatgpt.MessageCreateEventHandler)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Hi there!")
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
