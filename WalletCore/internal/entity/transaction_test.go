package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	clientA, _ := NewClient("Client A", "clientA@gmail.com")
	clientB, _ := NewClient("Client B", "clientB@gmail.com")

	accountFrom, _ := NewAccount(clientA)
	accountTo, _ := NewAccount(clientB)
	_ = accountFrom.Credit(100.0)
	transaction, _ := NewTransaction(accountFrom, accountTo, 100.0)
	assert.NotNil(t, transaction)
	assert.Equal(t, transaction.Amount, 100.0)
	assert.Equal(t, transaction.AccountFrom.Balance, 0.0)
	assert.Equal(t, transaction.AccountTo.Balance, 100.0)
}

func TestNewTransactionWithInvalidArgs(t *testing.T) {
	clientA, _ := NewClient("Client A", "clientA@gmail.com")
	clientB, _ := NewClient("Client B", "clientB@gmail.com")

	t.Run("amount must be greater than zero", func(t *testing.T) {
		accountFrom, _ := NewAccount(clientA)
		accountTo, _ := NewAccount(clientB)
		transaction, err := NewTransaction(accountFrom, accountTo, 0.0)
		assert.NotNil(t, err)
		assert.Nil(t, transaction)
	})

	t.Run("insufficient balance", func(t *testing.T) {
		accountFrom, _ := NewAccount(clientA)
		accountTo, _ := NewAccount(clientB)
		transaction, err := NewTransaction(accountFrom, accountTo, 100.0)
		assert.NotNil(t, err)
		assert.Nil(t, transaction)
	})

}
