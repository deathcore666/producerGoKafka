package main

import (
	"log"

	"github.com/pdevty/xvideos"
	"gopkg.in/telegram-bot-api.v4"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Topic: test"),
		tgbotapi.NewKeyboardButton("Topic: test1"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Topic: test2"),
		tgbotapi.NewKeyboardButton("Topic: test3"),
	),
)

var token string = "441748577:AAGb3HXzCSqw_jHUoar_CzBrKZzILTMb1ec"
var kafkaBrokers = []string{"localhost:9092"}

type xVid struct {
	url   string
	thumb string
}

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	//starting kafka client routine to listen to topic channnel
	var topicChan = make(chan string)
	var respChan = make(chan []byte)

	//bot
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			msgString := "Hello and welcome, " + update.Message.From.UserName + "!\n" +
				"/kafkasingletopic\n/kafkaall"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgString)
			bot.Send(msg)
		case "/kafkasingletopic":
			msgString := "Choose one"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgString)
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)
		case "/kafkaall":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
				"Reading  from all topics goroutinely"))

		case "Topic: test1":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
				"---------------------"))
			go consumer(topicChan, respChan, kafkaBrokers)
			topic := "test1"
			topicChan <- topic
			for msgStr := range respChan {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(msgStr))
				bot.Send(msg)
			}

		}

	}

	//myChan <- consumer.go
	/*
		myChan := make(chan []byte)
		topic := "test1"
		kafkaBrokers := []string{"localhost:9092"}
		go consumer(topic, myChan, kafkaBrokers)
	*/

}

func goGetit(quer string, xChan chan xVid) {
	xQuer := "http://jp.xvideos.com/c/" + quer + "/"
	xv, err := xvideos.Get(xQuer)

	if err != nil {
		log.Fatal(err)
	}

	var thumb, url string
	for _, v := range xv {
		thumb = v.Thumbnail
		url = v.Url
	}

	response := xVid{url, thumb}
	xChan <- response
}
