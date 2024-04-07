package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fabioods/fc-ms-wallet/internal/database"
	"github.com/fabioods/fc-ms-wallet/internal/event"
	"github.com/fabioods/fc-ms-wallet/internal/event/handler"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/create_account"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/create_client"
	"github.com/fabioods/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/fabioods/fc-ms-wallet/internal/web"
	"github.com/fabioods/fc-ms-wallet/internal/web/webserver"
	"github.com/fabioods/fc-ms-wallet/pkg/events"
	"github.com/fabioods/fc-ms-wallet/pkg/rabbitmq"
	"github.com/fabioods/fc-ms-wallet/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(mysql:3306)/wallet?parseTime=true"))
	if err != nil {
		fmt.Println("Error opening database: ", err)
		panic(err)
	}
	defer db.Close()

	// Verifica a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		panic(err)
	}

	connection := rabbitmq.ConnectToRabbitMQ("amqp://rabbitmq:rabbitmq@rabbitmq:5672/")

	exchangeTransactions := rabbitmq.NewExchange(connection, "direct", "transactions")
	err = exchangeTransactions.DeclareExchange()
	if err != nil {
		panic(err)
	}

	exchangeBalances := rabbitmq.NewExchange(connection, "direct", "balances")
	err = exchangeBalances.DeclareExchange()
	if err != nil {
		panic(err)
	}

	transactionsCreatedQueue := rabbitmq.NewQueue(connection, "transactions_created", "status:created", "transactions")
	err = transactionsCreatedQueue.DeclareQueue()
	if err != nil {
		panic(err)
	}

	balancesUpdatedQueue := rabbitmq.NewQueue(connection, "balances_updated", "status:updated", "balances")
	err = balancesUpdatedQueue.DeclareQueue()
	if err != nil {
		panic(err)
	}

	transactionProducer := rabbitmq.NewProducer(connection, "transactions_created", "status:created", "transactions")
	balanceProducer := rabbitmq.NewProducer(connection, "balance_updated", "status:updated", "balances")

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("transaction_created", handler.NewTransactionCreatedRabbitMQ(transactionProducer))
	eventDispatcher.Register("balance_updated", handler.NewBalanceUpdatedRabbitMQ(balanceProducer))

	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)

	ctx := context.Background()
	unitOfWork := uow.NewUow(ctx, db)

	unitOfWork.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	unitOfWork.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDB)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(unitOfWork, transactionCreatedEvent, balanceUpdatedEvent, eventDispatcher)

	webServer := webserver.NewWebServer("0.0.0.0:8080")

	clientHandler := web.NewClientHandlerWeb(*createClientUseCase)
	accountHandler := web.NewAccountHandlerWeb(*createAccountUseCase)
	transactionHandler := web.NewTransactionHandlerWeb(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClientHandlerWeb)
	webServer.AddHandler("/accounts", accountHandler.CreateAccountHandlerWeb)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransactionHandlerWeb)
	webServer.AddHandler("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	fmt.Println("Server started on port 8080")
	err = webServer.Start()
	if err != nil {
		panic(err)
	}
}
