package commands

import (
	"dizzle/config"
	"dizzle/utils"
	"fmt"
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

		if m.Content == "!kick" {
			kickMessage := "Kicked user: " + m.Author.Mention()

			s.ChannelMessageSend(m.ChannelID, kickMessage)
		}

		if m.Content == "!ban" {
			banMessage := "Banned user: " + m.Author.Mention()
			utils.AddBan(m.Author.Username)

			s.ChannelMessageSend(m.ChannelID, banMessage)
		}

		if m.Content == "!avatar" {
			reader, err := s.UserAvatarDecode(m.Member.User)

			if err != nil {
				fmt.Println(err.Error())
			}

			result := discordgo.MessageEmbed{
				Type:        "r",
				Title:       "Your avatar",
				Color:       7,
				Description: "Hello everyone",
				Image: &discordgo.MessageEmbedImage{
					URL: reader.Bounds().String(),
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, &result)

		}
	}
}
