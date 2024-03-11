package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Rabbit MQ")

	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Print(connection)
		panic(err)
	}

	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")
}
