package handlers

import (
	"net/http"
	"regexp"

	"github.com/felipeornelis/payment-simplified/application/errors"
	"github.com/felipeornelis/payment-simplified/application/validator"
	"github.com/felipeornelis/payment-simplified/infrastructure/handlers"
	"github.com/felipeornelis/payment-simplified/internal/account/interactors"
)

type CreateAccountHandler interface {
	handlers.Handler
}

type createAccountHandler struct {
	interactor interactors.CreateAccountInteractor
	handlers.BaseHandler
}

func NewCreateAccountHandler(i interactors.CreateAccountInteractor) CreateAccountHandler {
	return &createAccountHandler{
		interactor: i,
	}
}

// Handle implements CreateAccountHandler.
func (h *createAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var payload interactors.CreateAccountRequest
	if err := h.Decode(r.Body, &payload); err != nil {
		h.Error(w, err.Error(), err, http.StatusBadRequest)
		return
	}

	if err := validator.Validate(payload); err != nil {
		h.Error(w, "invalid request", err, http.StatusBadRequest)
		return
	}

	payload.Document = h.sanitizeDocumentNumber(payload.Document)

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

func (h *createAccountHandler) sanitizeDocumentNumber(document string) string {
	return regexp.MustCompile(`\D`).ReplaceAllString(document, "")
}
