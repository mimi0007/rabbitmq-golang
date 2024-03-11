package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Rabbit MQ")

	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	// The connection with rabbitmq will be closed within the execution of the main function
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	// This will open a unique channel for sending the bulk messages to the queue
	channel, err := connection.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer channel.Close()

	// Creating Queue to publish messages on that
	queue, err := channel.QueueDeclare(
		"TestRabbitMQ",
		false,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Queue: ", queue)
}
