package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "johndoe@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, client.Email, "johndoe@gmail.com")
}

func TestCreateNewClientWithInvalidArgs(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}
func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	err := client.Update("Fabio Santos", "fabio.santos@live.com")
	assert.Nil(t, err)
	assert.Equal(t, "Fabio Santos", client.Name)
	assert.Equal(t, "fabio.santos@live.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Fabio", "fabio@gmail.com")
	err := client.Update("", "")
	assert.NotNil(t, err)
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	account, _ := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
}

func TestAddAccountWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Fabio", "fah_ds@live.com")
	clientB, _ := NewClient("Fabio B", "fabioB@hotmail.com")
	account, _ := NewAccount(clientB)
	err := client.AddAccount(account)
	assert.NotNil(t, err)
}
