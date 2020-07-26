package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("go rabbitMQ Receiver")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to rabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"queuetest",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("received msgs: %s\n", d.Body)
		}
	}()

	fmt.Println("successufully connected to rabbitmq instance")
	fmt.Println(" [*] - waiting for msgs")
	<-forever

}
