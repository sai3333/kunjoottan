package main

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("899333580:AAH4BOeUZqTup-TTdJtNsjyQvCBipLuk69k")
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Print("Authorized on account %s")

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbot.NewMessage(update.Message.Chat.ID)
		var list := "fuck off",
		        "funda mayir",
		        "പോടീ പൂറി🤬🤬🤬🤬"
		msg := list
		bot.Send(list)
		
	}
}
