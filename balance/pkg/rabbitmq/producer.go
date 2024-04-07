package rabbitmq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Producer struct {
	Conn       *amqp.Connection
	RoutingKey string
	Exchange   string
	QueueName  string
}

func NewProducer(conn *amqp.Connection, queueName, routingKey, exchange string) *Producer {
	return &Producer{
		Conn:       conn,
		QueueName:  queueName,
		RoutingKey: routingKey,
		Exchange:   exchange,
	}
}

func (p Producer) Publisher(msg interface{}) error {
	ch, err := p.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = ch.Publish(
		p.Exchange,   // exchange
		p.RoutingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}
