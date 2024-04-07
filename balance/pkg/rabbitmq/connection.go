package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ(uri string) *amqp.Connection {
	conn, err := amqp.Dial(uri)
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to RabbitMQ: %s", err)
		panic(msg)
	}
	return conn
}
