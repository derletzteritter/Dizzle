package main

import (
	"dizzle/commands"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot" + "token")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Started Dizzle Bot: ", discord)

	discord.AddHandler(messageCreate)

	commands.CreateCommand("hello")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}
}
