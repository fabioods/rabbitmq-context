package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ(uri string) *amqp.Connection {
	conn, err := amqp.Dial(uri)
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to RabbitMQ: %s", err)
		panic(msg)
	}
	return conn
}
