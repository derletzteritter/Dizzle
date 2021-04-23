package bot

import (
	"dizzle/commands"
	"dizzle/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	discord, err := discordgo.New("Bot" + config.Token)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Started Dizzle Bot: ", discord)

	discord.AddHandler(commands.CreateCommand)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!ping" {
			s.ChannelMessageSend(m.ChannelID, "pong!")
		}
	}
}
