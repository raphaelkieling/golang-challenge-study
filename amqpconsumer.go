package main

import (
	"delivery-much-challenge/conf"
	"fmt"

	"github.com/streadway/amqp"
)

type AMQPConsumer struct {
	config *conf.Config
}

func (a *AMQPConsumer) Start() {
	dns := fmt.Sprintf("amqp://%s:%s@%s:5672/", a.config.AmqpUser, a.config.AmqpPass, a.config.AmqpHost)

	conn, err := amqp.Dial(dns)

	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "rabbitmq error to connect", err))
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to open a channel", err))
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("Test", false, false, false, false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to declare a queue", err))
	}
	defer ch.QueueDelete(queue.Name, false, false, true)

	err = ch.QueueBind(queue.Name, "#", "stock", false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to bind the queue", err))
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to consume", err))
	}

	fmt.Println("------------ START QUEUE")
	for msg := range msgs {
		fmt.Println("------------ FROM QUEUE")
		fmt.Println(msg)
	}
}
