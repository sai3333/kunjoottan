package main

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("987206364:AAGKPNcUnwPjXQdJOBUJ9RPjo6HZ735CHGY")
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

		msg1 := tgbot.NewMessage(update.Message.Chat.ID, "പോടീ പൂറി🤬🤬🤬🤬")
		
		
		bot.Send(msg1)
		
	}
}
