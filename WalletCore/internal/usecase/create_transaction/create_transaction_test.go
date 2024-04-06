package create_transaction

import (
	"context"
	"github.com/fabioods/fc-ms-wallet/internal/entity"
	"github.com/fabioods/fc-ms-wallet/internal/event"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/mocks"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateTransactionUseCase(t *testing.T) {
	client1, _ := entity.NewClient("Rafa", "rafa@example.com")
	client2, _ := entity.NewClient("Suzi", "suzie@example.com")
	account1, _ := entity.NewAccount(client1)
	account1.Credit(1000)
	account2, _ := entity.NewAccount(client2)
	account2.Credit(500)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDto{
		AccountIdFrom: account1.ID,
		AccountIdTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalanceUpdated := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, eventTransaction, eventBalanceUpdated, dispatcher)

	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
