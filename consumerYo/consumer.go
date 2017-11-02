package main

//var strCh = make(chan []byte)

<<<<<<< HEAD
import (
	"log"
=======
func consumer(reqChan chan kafkaRequest, respChan chan kafkaResponse, brokers []string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	select {
	case topic := <-reqChan:
		//get all partitions on the given topic
		partitionList, err := consumer.Partitions(topic.topic)
		if err != nil {
			fmt.Println("Error retrieving partitionList ", err)
		}

		//get offset for the oldest message on the topic --oldest-message
		initialOffset := sarama.OffsetOldest
		for _, partition := range partitionList {
			pc, _ := consumer.ConsumePartition(topic.topic, partition, initialOffset)
>>>>>>> b9f17810603af41506b908348ab9b337703530a5

	"gopkg.in/telegram-bot-api.v4"
)

<<<<<<< HEAD
func consumer(reqChan chan *tgbotapi.Message, respChan chan respCons) {
	str := "response text"
	for {
		for msg := range reqChan {
			respChan <- respCons{str, msg}
			respChan <- respCons{str, msg}
			log.Println("sent")
=======
			go func(pc sarama.PartitionConsumer) {
				for message := range pc.Messages() {
					respChan <- kafkaResponse{topic.telega, message.Value}
				}
			}(pc)
>>>>>>> b9f17810603af41506b908348ab9b337703530a5
		}
	}

}
