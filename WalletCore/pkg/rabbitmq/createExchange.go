package rabbitmq

import "github.com/streadway/amqp"

type Exchange struct {
	Conn         *amqp.Connection
	ExchangeType string
	ExchangeName string
}

func NewCreateExchange(connection *amqp.Connection, exchangeType, exchangeName string) *Exchange {
	return &Exchange{
		Conn:         connection,
		ExchangeType: exchangeType,
		ExchangeName: exchangeName,
	}
}

func (e Exchange) CreateExchange() error {
	channel, err := e.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare(
		e.ExchangeName, // nome
		e.ExchangeType, // tipo
		true,           // durável
		false,          // auto-delete
		false,          // interno
		false,          // no-wait
		nil,            // argumentos
	)
}
