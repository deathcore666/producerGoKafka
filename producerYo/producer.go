package main

import (
	"log"

	"github.com/gocql/gocql"

	"github.com/Shopify/sarama"
)

func newProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	return producer, err
}

func populate(session *gocql.Session) {
	for i := 1; i <= 20; i++ {
		if err := session.Query(`INSERT INTO users (id, lastname, name) VALUES (?, 'kuliev', 'bratan')`, i).Exec(); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}
