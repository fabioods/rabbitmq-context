package usecase

import "github.com/fabioods/balance/internal/entity"

type (
	ReportBalanceForAccountUseCase struct {
		AccountRepository AccountRepository
	}

	AccountRepository interface {
		FindByID(id string) (*entity.Account, error)
	}

	BalanceOutputDto struct {
		AccountID string  `json:"account_id"`
		Balance   float64 `json:"balance"`
	}
)

func NewReportBalanceForAccountUseCase(ar AccountRepository) *ReportBalanceForAccountUseCase {
	return &ReportBalanceForAccountUseCase{
		AccountRepository: ar,
	}
}

func (uc *ReportBalanceForAccountUseCase) Execute(id string) (BalanceOutputDto, error) {
	account, err := uc.AccountRepository.FindByID(id)
	if err != nil {
		return BalanceOutputDto{}, err
	}
	output := BalanceOutputDto{
		AccountID: account.ID,
		Balance:   account.Balance,
	}
	return output, nil
}
