package commands

import (
	"dizzle/config"
	"dizzle/functions"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CreateCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	functions.AvatarImage()

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "pong!")
		}
	}
}
