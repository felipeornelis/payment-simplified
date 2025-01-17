package interactors

import (
	"context"
	"net/http"

	"github.com/felipeornelis/payment-simplified/application/errors"
	"github.com/felipeornelis/payment-simplified/internal/account"
	"github.com/felipeornelis/payment-simplified/internal/account/mappers"
)

type CreateAccountInteractor interface {
	Execute(context.Context, CreateAccountRequest) (mappers.AccountDTO, error)
}

type CreateAccountRequest struct {
	Name     string  `json:"name" validate:"required"`
	Document string  `json:"document" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Type     string  `json:"type" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Balance  float64 `json:"balance" validate:"required"`
}

type createAccountInteractor struct {
	repository account.Repository
}

func NewCreateAccountInteractor(r account.Repository) CreateAccountInteractor {
	return &createAccountInteractor{
		repository: r,
	}
}

// Execute implements CreateAccountInteractor.
func (i *createAccountInteractor) Execute(ctx context.Context, payload CreateAccountRequest) (mappers.AccountDTO, error) {
	_, err := i.repository.FindByDocument(ctx, payload.Document)
	if err == nil {
		if len(payload.Document) == 11 {
			return mappers.AccountDTO{}, errors.New(http.StatusConflict, account.ErrCPFRegistered)
		}

		return mappers.AccountDTO{}, errors.New(http.StatusConflict, account.ErrCNPJRegistered)
	}

	_, err = i.repository.FindByEmail(ctx, payload.Email)
	if err == nil {
		return mappers.AccountDTO{}, errors.New(http.StatusConflict, account.ErrEmailAddressRegistered)
	}

	accountType, err := account.ParseType(payload.Type)
	if err != nil {
		return mappers.AccountDTO{}, errors.New(http.StatusBadRequest, account.ErrUnknownAccountType)
	}

	// TODO: hash password before passing it to constructor
	acc, err := account.New(payload.Name, payload.Email, payload.Password, payload.Document, payload.Balance, accountType)
	if err != nil {
		return mappers.AccountDTO{}, errors.New(http.StatusBadRequest, err)
	}

	if err = i.repository.Save(ctx, acc); err != nil {
		return mappers.AccountDTO{}, errors.New(http.StatusInternalServerError, err)
	}

	return mappers.AccountToDTO(acc), nil
}
