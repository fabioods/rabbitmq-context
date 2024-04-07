package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	Conn      *amqp.Connection
	QueueName string
}

func NewConsumer(conn *amqp.Connection, queueName string) *Consumer {
	return &Consumer{
		Conn:      conn,
		QueueName: queueName,
	}
}

func (c Consumer) Consume(msgChan chan amqp.Delivery) {
	ch, err := c.Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(c.QueueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		msgChan <- d
	}
}
