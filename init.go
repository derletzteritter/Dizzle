package main

import (
	"dizzle/bot"
	"dizzle/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/* database.Start() */
	bot.Start()
}
