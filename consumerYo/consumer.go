package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//var strCh = make(chan []byte)

//Consumer kafka takes topic as 1 arg
func consumer(topic string, strCh chan []byte, brokers []string) {
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

	//get all partitions on the given topic
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}

	//get offset for the oldest message on the topic --oldest-message
	initialOffset := sarama.OffsetOldest
	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				strCh <- message.Value
			}
		}(pc)
	}
}
