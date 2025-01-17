package main

import (
	"net/http"

	"github.com/felipeornelis/payment-simplified/internal/account/handlers"
	"github.com/felipeornelis/payment-simplified/internal/account/interactors"
	"github.com/felipeornelis/payment-simplified/internal/account/persistence"

	txhandlers "github.com/felipeornelis/payment-simplified/internal/transaction/handlers"
	txinteractors "github.com/felipeornelis/payment-simplified/internal/transaction/interactors"
	txpersistence "github.com/felipeornelis/payment-simplified/internal/transaction/persistence"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	accountRepository := persistence.NewMemoryAccountRepository()
	createAccountInteractor := interactors.NewCreateAccountInteractor(accountRepository)
	createAccountHandler := handlers.NewCreateAccountHandler(createAccountInteractor)
	listAccountsInteractor := interactors.NewListAccountsInteractor(accountRepository)
	listAccountsHandler := handlers.NewListAccountsHandler(listAccountsInteractor)

	transactionRepository := txpersistence.NewTransactionRepository()
	createTransactionInteractor := txinteractors.NewCreateTransactionInteractor(transactionRepository, accountRepository)
	createTransactionHandler := txhandlers.NewCreateTransactionHandler(createTransactionInteractor)

	r.Post("/accounts", createAccountHandler.Handle)
	r.Get("/accounts", listAccountsHandler.Handle)
	r.Post("/transfer", createTransactionHandler.Handle)

	http.ListenAndServe(":8080", r)
}
