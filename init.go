package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot" + "token")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Started Dizzle Bot: ", discord)
}
