package event

import "time"

type TransactionCreated struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name: "transaction_created",
	}
}

func (t *TransactionCreated) GetName() string {
	return t.Name
}

func (t *TransactionCreated) SetPayload(payload interface{}) {
	t.Payload = payload
}

func (t *TransactionCreated) GetPayload() interface{} {
	return t.Payload
}

func (t *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}
