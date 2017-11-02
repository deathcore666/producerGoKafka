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

<<<<<<< HEAD
type respCons struct {
	resp string
	msg  *tgbotapi.Message
=======
type kafkaResponse struct {
	telega  *tgbotapi.Message
	message []byte
}

type kafkaRequest struct {
	telega *tgbotapi.Message
	topic  string
>>>>>>> b9f17810603af41506b908348ab9b337703530a5
}

func main() {
	//channels for request response
	var reqChan = make(chan kafkaRequest)
	var respChan = make(chan kafkaResponse)

	//starting kafka client routine to listen to topic channnel
	go consumer(reqChan, respChan, kafkaBrokers)

	//bot thingy here
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
<<<<<<< HEAD
	//starting kafka client routine to listen to topic channnel
	var respChan = make(chan respCons)
	var reqChan = make(chan *tgbotapi.Message)
	go consumer(reqChan, respChan)
	//bot
=======
>>>>>>> b9f17810603af41506b908348ab9b337703530a5
	for {
		select {
		case update := <-updates:
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
				numericKeyboard.OneTimeKeyboard = true
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			case "/kafkaall":
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"Reading  from all topics goroutinely"))

			case "Topic: test1":
<<<<<<< HEAD
				reqChan <- update.Message
			}
		case resp := <-respChan:
			bot.Send(tgbotapi.NewMessage(resp.msg.Chat.ID, resp.resp))
=======
				topic := "test1"
				reqChan <- kafkaRequest{update.Message, topic}
			}
		case response := <-respChan:
			bot.Send(tgbotapi.NewMessage(response.telega.Chat.ID, string(response.message)))
>>>>>>> b9f17810603af41506b908348ab9b337703530a5
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
