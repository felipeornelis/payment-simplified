package handlers

import (
	"net/http"

	"github.com/felipeornelis/payment-simplified/application/errors"
	"github.com/felipeornelis/payment-simplified/application/validator"
	"github.com/felipeornelis/payment-simplified/infrastructure/handlers"
	"github.com/felipeornelis/payment-simplified/internal/transaction/interactors"
)

type CreateTransactionHandler interface {
	handlers.Handler
}

type createTransactionHandler struct {
	interactor interactors.CreateTransactionInteractor
	handlers.BaseHandler
}

func NewCreateTransactionHandler(i interactors.CreateTransactionInteractor) CreateTransactionHandler {
	return &createTransactionHandler{
		interactor: i,
	}
}

// Handle implements CreateTransactionHandler.
func (h *createTransactionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var payload interactors.CreateTransactionRequest
	if err := h.Decode(r.Body, &payload); err != nil {
		h.Error(w, err.Error(), err, http.StatusBadRequest)
		return
	}

	if err := validator.Validate(payload); err != nil {
		h.Error(w, "invalid request", err, http.StatusBadRequest)
		return
	}

	output, err := h.interactor.Execute(r.Context(), payload)
	if err != nil {
		outputErr := err.(errors.AppError)
		if outputErr.Code != http.StatusInternalServerError {
			h.Error(w, outputErr.Message, outputErr.Err, outputErr.Code)
			return
		}

		h.Error(w, "unexpected error", outputErr.Err, outputErr.Code)
		return
	}

	h.Response(w, output, http.StatusCreated)
}
