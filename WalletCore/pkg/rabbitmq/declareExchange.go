package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Exchange struct {
	Conn         *amqp.Connection
	ExchangeType string
	ExchangeName string
}

func NewExchange(connection *amqp.Connection, exchangeType, exchangeName string) *Exchange {
	return &Exchange{
		Conn:         connection,
		ExchangeType: exchangeType,
		ExchangeName: exchangeName,
	}
}

func (e Exchange) DeclareExchange() error {
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
