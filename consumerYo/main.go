package main

import (
	"log"

	"github.com/pdevty/xvideos"
	"gopkg.in/telegram-bot-api.v4"
)

type xVid struct {
	url   string
	thumb string
}

func main() {
	bot, err := tgbotapi.NewBotAPI("441748577:AAGb3HXzCSqw_jHUoar_CzBrKZzILTMb1ec")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	xReq := make(chan string)
	xRes := make(chan xVid)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//updates, err := bot.GetUpdatesChan(u)

	//myChan <- consumer.go
	/*
		myChan := make(chan []byte)
		topic := "test1"
		kafkaBrokers := []string{"localhost:9092"}
		go consumer(topic, myChan, kafkaBrokers)
	*/

	helloBool := false
	searchBool := false
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	botChan, err := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-botChan:
			if helloBool == false {
				hi := "Hello, " + update.Message.From.UserName + ". I am xvid bot" + "\n" +
					"I will show you the world. Type /search to start our journey!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, hi)
				bot.Send(msg)
				helloBool = true
			}
			if searchBool == true {
				xReq <- update.Message.Text
			}
			if update.Message.Text == "/search" {
				searchBool = true
			}

		case resp := <-xRes:

		}
	}
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
