package main

//var strCh = make(chan []byte)

import "gopkg.in/telegram-bot-api.v4"

func consumer(reqChan chan *tgbotapi.Message, respChan chan respCons) {
	str := "response text"
	for {
		for msg := range reqChan {
			respChan <- respCons{str, msg}
		}
	}
}
