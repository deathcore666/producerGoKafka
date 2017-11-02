package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//var strCh = make(chan []byte)

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
	case request := <-reqChan:
		//get all partitions on the given topic
		partitionList, err := consumer.Partitions(request.topic)
		if err != nil {
			fmt.Println("Error retrieving partitionList ", err)
		}

		initialOffset := sarama.OffsetOldest
		for _, partition := range partitionList {
			pc, _ := consumer.ConsumePartition(request.topic, partition, initialOffset)
			
			go func(pc sarama.PartitionConsumer) {
				for {
					select {
					case message := <-pc.Messages():
						respChan <- kafkaResponse{request.telega, message.Value}

					}
				}
			}(pc)
		}
	}
}
