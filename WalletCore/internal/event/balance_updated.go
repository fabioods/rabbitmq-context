package event

import "time"

type BalanceUpdated struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

func NewBalanceUpdated() *BalanceUpdated {
	return &BalanceUpdated{
		Name: "balance_updated",
	}
}

func (t *BalanceUpdated) GetName() string {
	return t.Name
}

func (t *BalanceUpdated) SetPayload(payload interface{}) {
	t.Payload = payload
}

func (t *BalanceUpdated) GetPayload() interface{} {
	return t.Payload
}

func (t *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
