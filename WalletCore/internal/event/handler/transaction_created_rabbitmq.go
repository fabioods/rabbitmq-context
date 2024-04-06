package handler

import (
	"fmt"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/rabbitmq"
	"sync"
)

type TransactionCreatedRabbitMQHandler struct {
	RabbitMQ *rabbitmq.Producer
}

func NewTransactionCreatedRabbitMQ(rabbitmq *rabbitmq.Producer) *TransactionCreatedRabbitMQHandler {
	return &TransactionCreatedRabbitMQHandler{
		RabbitMQ: rabbitmq,
	}
}

func (t *TransactionCreatedRabbitMQHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	err := t.RabbitMQ.Publisher(message)
	if err != nil {
		fmt.Println("Error to publish message to rabbitmq", err)
		return
	}
	fmt.Println("Publishing message to rabbitmq", message.GetPayload())
}
