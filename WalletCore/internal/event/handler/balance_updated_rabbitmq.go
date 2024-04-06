package handler

import (
	"fmt"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/rabbitmq"
	"sync"
)

type BalanceUpdatedRabbitMQHandler struct {
	RabbitMQ *rabbitmq.Producer
}

func NewBalanceUpdatedRabbitMQ(rabbitmq *rabbitmq.Producer) *BalanceUpdatedRabbitMQHandler {
	return &BalanceUpdatedRabbitMQHandler{
		RabbitMQ: rabbitmq,
	}
}

func (t *BalanceUpdatedRabbitMQHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	err := t.RabbitMQ.Publisher(event)
	if err != nil {
		fmt.Sprintf("Error to publish message to RabbitMQ ", err)
		return
	}
	fmt.Println("Publishing message to rabbitMQ ", event.GetPayload())
}
