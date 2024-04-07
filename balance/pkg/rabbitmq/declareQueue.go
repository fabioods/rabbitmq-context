package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Queue struct {
	Conn       *amqp.Connection
	QueueName  string
	RoutingKey string
	Exchange   string
}

func NewQueue(conn *amqp.Connection, queueName, routingKey, exchange string) *Queue {
	return &Queue{
		Conn:       conn,
		QueueName:  queueName,
		RoutingKey: routingKey,
		Exchange:   exchange,
	}
}

func (q Queue) CreateQueue() error {
	ch, err := q.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	queueDeclare, err := ch.QueueDeclare(q.QueueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	return ch.QueueBind(queueDeclare.Name, q.RoutingKey, q.Exchange, false, nil)
}
