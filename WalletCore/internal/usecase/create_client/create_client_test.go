package create_client

import (
	"github.com/fabioods/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (cg *ClientGatewayMock) Save(client *entity.Client) error {
	args := cg.Called(client)
	return args.Error(0)
}

func (cg *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := cg.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestNewCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)
	out, err := uc.Execute(InputDTO{
		Name:  "Fabio",
		Email: "fabio@gmail.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Fabio", out.Name)
	m.AssertExpectations(t)
	m.AssertCalled(t, "Save", mock.Anything)
	m.AssertNumberOfCalls(t, "Save", 1)
}
