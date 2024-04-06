package handler

import (
	"fmt"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/kafka"
	"sync"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafka(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (t *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	err := t.Kafka.Publish(message, nil, "transactions")
	if err != nil {
		fmt.Println("Error to publish message to kafka", err)
		return
	}
	fmt.Println("Publishing message to kafka", message.GetPayload())
}
