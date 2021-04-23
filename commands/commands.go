package commands

import (
	"dizzle/config"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CreateCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "pong!")
		}

		if m.Content == "!warn" {

			warnMessage := "Warned user: " + m.Author.Mention()

			s.ChannelMessageSend(m.ChannelID, warnMessage)
		}

		if m.Content == "!avatar" {
			result := discordgo.MessageEmbed{
				Type:        "image",
				Title:       "Your avatar",
				Color:       7,
				Description: "Hello everyone",
				Video: &discordgo.MessageEmbedVideo{
					URL:    "https://www.youtube.com/watch?v=somevideo",
					Width:  50,
					Height: 50,
				},
				Author: &discordgo.MessageEmbedAuthor{
					Name: m.Author.Username,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, &result)

		}
	}
}
