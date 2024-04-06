package create_transaction

import (
	"context"
	"fmt"
	"github.com/fabioods/fc-ms-wallet/internal/entity"
	"github.com/fabioods/fc-ms-wallet/internal/gateway"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/uow"
)

type CreateTransactionInputDto struct {
	AccountIdFrom string  `json:"account_id_from"`
	AccountIdTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	Id          string  `json:"id"`
	AccountFrom string  `json:"account_from"`
	AccountTo   string  `json:"account_to"`
	Amount      float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

// variaveis para DI
type CreateTransactionUseCase struct {
	UnitOfWork         uow.UowInterface
	TransactionCreated events.EventInterface
	BalanceUpdated     events.EventInterface
	EventsDispatcher   events.EventDispatcherInterface
}

func NewCreateTransactionUseCase(uow uow.UowInterface, tc events.EventInterface, bu events.EventInterface, ed events.EventDispatcherInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		UnitOfWork:         uow,
		TransactionCreated: tc,
		EventsDispatcher:   ed,
		BalanceUpdated:     bu,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	output := &CreateTransactionOutputDto{}
	balanceUpdatedOutputDTO := BalanceUpdatedOutputDTO{}
	err := uc.UnitOfWork.Do(ctx, func(uow *uow.Uow) error {
		accountFrom, err := uc.getAccountRepository(ctx).FindByID(input.AccountIdFrom)
		if err != nil {
			fmt.Println("Error to find account from")
			return err
		}
		accountTo, err := uc.getAccountRepository(ctx).FindByID(input.AccountIdTo)
		if err != nil {
			fmt.Println("Error to find account to")
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

		if err != nil {
			fmt.Println("Error to create transaction")
			return err
		}

		err = uc.getAccountRepository(ctx).UpdateBalance(accountFrom)
		if err != nil {
			fmt.Println("Error to update account from")
			return err
		}

		err = uc.getAccountRepository(ctx).UpdateBalance(accountTo)
		if err != nil {
			fmt.Println("Error to update account to")
			return err
		}

		err = uc.getTransactionRepository(ctx).Create(transaction)

		if err != nil {
			fmt.Println("Error to execute transaction")
			return err
		}
		output.Id = transaction.ID
		output.AccountFrom = transaction.AccountFrom.ID
		output.AccountTo = transaction.AccountTo.ID
		output.Amount = transaction.Amount

		balanceUpdatedOutputDTO.AccountIDFrom = input.AccountIdFrom
		balanceUpdatedOutputDTO.AccountIDTo = input.AccountIdTo
		balanceUpdatedOutputDTO.BalanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutputDTO.BalanceAccountIDTo = accountTo.Balance
		return nil

	})
	uc.TransactionCreated.SetPayload(output)
	err = uc.EventsDispatcher.Dispatch(uc.TransactionCreated)
	if err != nil {
		fmt.Println("Error to dispatch transaction created")
		return nil, err
	}

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutputDTO)
	err = uc.EventsDispatcher.Dispatch(uc.BalanceUpdated)
	if err != nil {
		fmt.Println("Error to dispatch balance updated")
		return nil, err

	}

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.UnitOfWork.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.UnitOfWork.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
