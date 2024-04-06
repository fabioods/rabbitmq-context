package gateway

import "github.com/fabioods/balance/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
