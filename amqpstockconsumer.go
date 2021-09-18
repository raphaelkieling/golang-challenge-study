package main

import (
	"delivery-much-challenge/conf"
	"fmt"

	"github.com/streadway/amqp"
)

type AMQPStockConsumer struct {
	config *conf.Config
}

func (a *AMQPStockConsumer) Start() chan<- string {
	dns := fmt.Sprintf("amqp://%s:%s@%s:5672/", a.config.AmqpUser, a.config.AmqpPass, a.config.AmqpHost)

	conn, err := amqp.Dial(dns)

	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "rabbitmq error to connect", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to open a channel", err))
	}

	queue, err := ch.QueueDeclare("stock-consumer", false, false, false, false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to declare a queue", err))
	}

	err = ch.QueueBind(queue.Name, "incremented", "stock", false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to bind the incremented", err))
	}

	err = ch.QueueBind(queue.Name, "decremented", "stock", false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to bind the decremented", err))
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %s", "failed to consume", err))
	}

	stockChannel := make(chan string)

	go func() {
		for msg := range msgs {
			stockChannel <- string(msg.Body)
		}
	}()

	return stockChannel
}
