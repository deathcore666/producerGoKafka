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

	//myChan <- consumer.go
	myChan := make(chan []byte)
	topic := "test1"
	kafkaBrokers := []string{"localhost:9092"}
	go consumer(topic, myChan, kafkaBrokers)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//myChan <- consumer.go
		for msgStr := range myChan {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(msgStr))
			bot.Send(msg)
		}

	}

}
