package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("441748577:AAGb3HXzCSqw_jHUoar_CzBrKZzILTMb1ec")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	topic := "test1"
	go consumer(topic)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//strCh comes from the consumer goroutine
		for msgStr := range strCh {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(msgStr))
			bot.Send(msg)
		}

	}

}
