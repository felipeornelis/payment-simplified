package handlers

import (
	"net/http"

	"github.com/felipeornelis/payment-simplified/infrastructure/handlers"
	"github.com/felipeornelis/payment-simplified/internal/account/interactors"
)

type ListAccountsHandler interface {
	handlers.Handler
}

type listAccountsHandler struct {
	interactor interactors.ListAccountsInteractor
	handlers.BaseHandler
}

func NewListAccountsHandler(i interactors.ListAccountsInteractor) ListAccountsHandler {
	return &listAccountsHandler{
		interactor: i,
	}
}

// Handle implements ListAccountsHandler.
func (h *listAccountsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	output := h.interactor.Execute(r.Context())
	h.Response(w, output, http.StatusOK)
}
