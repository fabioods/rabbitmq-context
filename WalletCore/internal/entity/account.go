package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	Client    *Client   `json:"client"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("client is required")
	}

	return nil
}

func (a *Account) Credit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	a.Balance += amount
	a.UpdateAt = time.Now()
	return nil
}

func (a *Account) Debit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	a.UpdateAt = time.Now()
	return nil
}
