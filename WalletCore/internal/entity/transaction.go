package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID          string    `json:"id"`
	AccountFrom *Account  `json:"account_from"`
	AccountTo   *Account  `json:"account_to"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	t := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, err
	}
	err = t.Execute()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Transaction) Execute() error {
	err := t.AccountFrom.Debit(t.Amount)
	if err != nil {
		return err
	}
	err = t.AccountTo.Credit(t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient balance")
	}
	return nil
}
