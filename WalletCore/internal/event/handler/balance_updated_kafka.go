package handler

import (
	"fmt"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/kafka"
	"sync"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafka(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{
		Kafka: kafka,
	}
}

func (t *BalanceUpdatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	err := t.Kafka.Publish(message, nil, "balances")
	if err != nil {
		fmt.Println("Error to publish message to kafka", err)
		return
	}
	fmt.Println("Publishing message to kafka", message.GetPayload())
}
