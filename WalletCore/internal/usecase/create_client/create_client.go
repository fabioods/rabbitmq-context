package create_client

import (
	"fmt"
	"github.com/fabioods/fc-ms-wallet/internal/entity"
	"github.com/fabioods/fc-ms-wallet/internal/gateway"
	"time"
)

type (
	InputDTO struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	OutputDTO struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UseCase struct {
		ClientGateway gateway.ClientGateway
	}
)

func (uc *UseCase) Execute(input InputDTO) (*OutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &OutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil

}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *UseCase {
	return &UseCase{
		ClientGateway: clientGateway,
	}
}
