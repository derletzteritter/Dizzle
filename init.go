package main

import (
	"dizzle/bot"
	"dizzle/config"
	"dizzle/database"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	database.Start()
	bot.Start()
}
