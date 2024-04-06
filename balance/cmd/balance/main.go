package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fabioods/balance/internal/usecase"
	"github.com/fabioods/balance/internal/web/webserver"
	"log"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabioods/balance/internal/database"
	"github.com/fabioods/balance/pkg/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(mysql:3306)/balances?parseTime=true"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	accountDB := database.NewAccountDB(db)
	usecaseProcessTransaction := usecase.NewProcessTransactionUseCase(accountDB)
	findusecase := usecase.NewFindByIDUseCase(accountDB)
	reportBalanceUseCase := usecase.NewReportBalanceForAccountUseCase(accountDB)

	consumer := initKafka()
	msgChan := make(chan *ckafka.Message)
	go processMessage(ctx, msgChan, usecaseProcessTransaction)
	go consume(ctx, consumer, msgChan)

	webServer := webserver.NewWebServer("0.0.0.0:3003")
	webServer.AddHandler("/ping", pingHandler)
	webServer.AddHandler("/accounts/{id}", findAccountHandler(findusecase))
	webServer.AddHandler("/balances/{id}", reportBalanceHandler(reportBalanceUseCase))

	log.Println("Server started on port 3003")
	if err := webServer.Start(); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}

func initKafka() *kafka.Consumer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	}

	topics := []string{"transactions"}
	return kafka.NewConsumer(configMap, topics)
}

func processMessage(ctx context.Context, msgChan chan *ckafka.Message, uc *usecase.ProcessTransactionUseCase) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgChan:
			fmt.Printf("Message Received: %s\n", string(msg.Value))
			var transactionDTO usecase.ProcessTransactionInputDto
			if err := json.Unmarshal(msg.Value, &transactionDTO); err != nil {
				log.Fatalf("Error to unmarshal message: %v", err)
			}
			if err := uc.Execute(transactionDTO); err != nil {
				log.Printf("Error processing transaction: %v", err)
			}
		}
	}
}

func consume(ctx context.Context, consumer *kafka.Consumer, msgChan chan *ckafka.Message) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := consumer.Consume(msgChan)
			if err != nil {
				log.Fatalf("Failed to consume messages: %v", err)
			}
		}
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func findAccountHandler(uc *usecase.FindByIDUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is required"))
			return
		}
		dto := usecase.FindByIDInputDto{
			ID: id,
		}
		account, err := uc.Execute(dto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(account)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func reportBalanceHandler(reportBalanceUseCase *usecase.ReportBalanceForAccountUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is required"))
			return
		}
		balance, err := reportBalanceUseCase.Execute(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		jsonData, _ := json.Marshal(balance)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

}
