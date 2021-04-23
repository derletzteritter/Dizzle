package bot

import (
	"dizzle/commands"
	"dizzle/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	dg, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Started Dizzle Bot")

	dg.AddHandler(commands.CreateCommand)

	<-make(chan struct{})
}
