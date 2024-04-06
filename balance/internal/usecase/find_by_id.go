package usecase

import (
	"github.com/fabioods/balance/internal/entity"
	"github.com/fabioods/balance/internal/gateway"
)

type (
	FindByIDUseCase struct {
		AccountRepository gateway.AccountGateway
	}

	FindByIDInputDto struct {
		ID string `json:"id"`
	}
)

func NewFindByIDUseCase(ar gateway.AccountGateway) *FindByIDUseCase {
	return &FindByIDUseCase{
		AccountRepository: ar,
	}
}

func (uc *FindByIDUseCase) Execute(input FindByIDInputDto) (*entity.Account, error) {
	return uc.AccountRepository.FindByID(input.ID)
}
