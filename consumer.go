package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Consumer")

	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer connection.Close()

	channel, err := connection.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer channel.Close()

	msgs, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	newChannel := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to Rabbitmq Instance\nwaiting for messages......")

	<-newChannel

}
