package web

import (
	"encoding/json"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/create_account"
	"net/http"
)

type AccountHandlerWeb struct {
	CreateAccountUseCase create_account.CreateAccountUseCase
}

func NewAccountHandlerWeb(createAccountUseCase create_account.CreateAccountUseCase) *AccountHandlerWeb {
	return &AccountHandlerWeb{
		CreateAccountUseCase: createAccountUseCase,
	}
}

func (a *AccountHandlerWeb) CreateAccountHandlerWeb(w http.ResponseWriter, r *http.Request) {
	var dto create_account.CreateAccountInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := a.CreateAccountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
