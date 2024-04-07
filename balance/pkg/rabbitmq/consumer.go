package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
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

func (c Consumer) Consume(msgChan chan amqp.Delivery) error {
	ch, err := c.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(c.QueueName, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to register a consumer: %s", err)
		return err
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		msgChan <- d
	}
	return nil
}
