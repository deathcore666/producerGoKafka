package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
)

var brokers = []string{"localhost:9092"}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	//yuolo
	return msg
}

func main() {

	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}
	topic := "test1"

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demodb"
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//db population used once!
	/*
		for i := 1; i <= 20; i++ {
			if err := session.Query(`INSERT INTO users (id, lastname, name) VALUES (?, 'kuliev', 'bratan')`, i).Exec(); err != nil {
				log.Fatal("Error: ", err)
			}
		}
	*/
	var msg string
	var name, lastname string
	query := `SELECT name,lastname FROM users`
	iter := session.Query(query).Iter()
	for iter.Scan(&name, &lastname) {
		msg = name + " " + lastname
		message := prepareMessage(topic, msg)
		producer.SendMessage(message)
	}
}
