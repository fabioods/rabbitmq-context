package web

import (
	"encoding/json"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/create_client"
	"net/http"
)

type ClientHandlerWeb struct {
	CreateClientUseCase create_client.UseCase
}

func NewClientHandlerWeb(createClientUseCase create_client.UseCase) *ClientHandlerWeb {
	return &ClientHandlerWeb{
		CreateClientUseCase: createClientUseCase,
	}
}

func (c *ClientHandlerWeb) CreateClientHandlerWeb(w http.ResponseWriter, r *http.Request) {
	var dto create_client.InputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := c.CreateClientUseCase.Execute(dto)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
