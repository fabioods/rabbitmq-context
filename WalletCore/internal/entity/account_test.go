package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestCreateNewAccountWithInvalidArgs(t *testing.T) {
	account, err := NewAccount(nil)
	assert.NotNil(t, err)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	err := account.Credit(100)
	assert.Nil(t, err)
}

func TestCreditAccountWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	err := account.Credit(0)
	assert.NotNil(t, err)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	_ = account.Credit(100)
	err := account.Debit(50)
	assert.Nil(t, err)
}

func TestDebitAccountWithZeroArg(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	_ = account.Credit(100)
	err := account.Debit(0)
	assert.NotNil(t, err)
}

func TestDebitAccountWithDebitHigherThanBalance(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	_ = account.Credit(10)
	err := account.Debit(20)
	assert.NotNil(t, err)
}
