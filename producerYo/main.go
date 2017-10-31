package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	//Kafka
	var brokers = []string{"localhost:9092"}
	producer, err := newProducer(brokers)
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}
	topic := "test1"

	//Cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demodb"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//------------------------
	//db population used once!
	//populate(session)
	//------------------------

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
