package usecase

import (
	"context"
	"github.com/fabioods/balance/internal/gateway"
	"github.com/fabioods/balance/pkg/rollback"
	"log"
)

type (
	ProcessTransactionUseCase struct {
		AccountRepository gateway.AccountGateway
	}

	ProcessTransactionInputDto struct {
		Name    string  `json:"name"`
		Payload payload `json:"payload"`
	}

	payload struct {
		ID          string  `json:"id"`
		AccountFrom string  `json:"account_from"`
		AccountTo   string  `json:"account_to"`
		Amount      float64 `json:"amount"`
	}
)

func NewProcessTransactionUseCase(ar gateway.AccountGateway) *ProcessTransactionUseCase {
	return &ProcessTransactionUseCase{
		AccountRepository: ar,
	}
}

func (uc *ProcessTransactionUseCase) Execute(input ProcessTransactionInputDto) error {
	rb := rollback.New()
	ctx := context.Background()

	accountFrom, err := uc.AccountRepository.FindByID(input.Payload.AccountFrom)
	if err != nil {
		return err
	}

	accountTo, err := uc.AccountRepository.FindByID(input.Payload.AccountTo)
	if err != nil {
		return err
	}

	err = accountFrom.Debit(input.Payload.Amount)
	if err != nil {
		return err
	}

	rb.Add("Debit accountFrom", func() {
		err := accountFrom.Credit(input.Payload.Amount)
		if err != nil {
			log.Println("Error to credit accountFrom in rollback")
		}
		err = uc.AccountRepository.UpdateBalance(accountFrom)
		if err != nil {
			log.Println("Error to update accountFrom in rollback")
		}
	})

	err = uc.AccountRepository.UpdateBalance(accountFrom)
	if err != nil {
		return err
	}

	err = accountTo.Credit(input.Payload.Amount)
	if err != nil {
		_ = rb.Do(ctx)
		return err
	}

	err = uc.AccountRepository.UpdateBalance(accountTo)
	if err != nil {
		_ = rb.Do(ctx)
		return err
	}

	return nil
}
