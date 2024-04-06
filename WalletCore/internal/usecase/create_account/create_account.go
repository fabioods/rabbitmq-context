package create_account

import (
	"github.com/fabioods/fc-ms-wallet/internal/entity"
	"github.com/fabioods/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientId string `json:"client_id"`
}

type CreateAccountOutputDto struct {
	Id string `json:"id"`
}

// variaveis da injeção de dependencia
type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

// construtor
func NewCreateAccountUseCase(ag gateway.AccountGateway, cg gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: ag,
		ClientGateway:  cg,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := uc.ClientGateway.Get(input.ClientId)

	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}

	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	output := &CreateAccountOutputDto{
		Id: account.ID,
	}
	return output, nil

}
