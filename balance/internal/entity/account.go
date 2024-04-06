package entity

import (
	"errors"
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	ClientID  string    `json:"client_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func NewAccount(id, clientID string, balance float64) (*Account, error) {
	account := &Account{
		ID:       id,
		ClientID: clientID,
		Balance:  balance,
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Validate() error {
	if a.ID == "" {
		return errors.New("id is required")
	}

	if a.ClientID == "" {
		return errors.New("client is required")
	}

	if a.Balance < 0 {
		return errors.New("balance must be greater than or equal to zero")
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
