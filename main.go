package main

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("1272367761:AAEWiHnT4QEtaq8xKLZppG5TG1DVj2BkRIM") //token from botfather
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Print("connected duccessfully")

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		
	
		msg := tgbot.NewMessage(update.Message.Chat.ID, "Hey Must join t.me/The_Tech_Hackers") //message
		bot.Send(msg)
	}
}
