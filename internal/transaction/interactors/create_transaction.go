package interactors

import (
	"context"
	"net/http"

	"github.com/felipeornelis/payment-simplified/application/errors"
	"github.com/felipeornelis/payment-simplified/internal/account"
	"github.com/felipeornelis/payment-simplified/internal/transaction"
	"github.com/felipeornelis/payment-simplified/internal/transaction/mappers"
)

type CreateTransactionInteractor interface {
	Execute(context.Context, CreateTransactionRequest) (mappers.TransactionDTO, error)
}

type CreateTransactionRequest struct {
	Value float64 `json:"value" validate:"required"`
	Payer string  `json:"payer" validate:"required"`
	Payee string  `json:"payee" validate:"required"`
}

type createTransactionInteractor struct {
	repository        transaction.Repository
	accountRepository account.Repository
}

func NewCreateTransactionInteractor(r transaction.Repository, ar account.Repository) CreateTransactionInteractor {
	return &createTransactionInteractor{
		repository:        r,
		accountRepository: ar,
	}
}

// Execute implements CreateTransactionInteractor.
func (i *createTransactionInteractor) Execute(ctx context.Context, payload CreateTransactionRequest) (mappers.TransactionDTO, error) {
	payer, err := i.accountRepository.FindByID(ctx, payload.Payer)
	if err != nil {
		return mappers.TransactionDTO{}, errors.New(http.StatusNotFound, account.ErrNotFound)
	}

	payee, err := i.accountRepository.FindByID(ctx, payload.Payee)
	if err != nil {
		return mappers.TransactionDTO{}, errors.New(http.StatusNotFound, account.ErrNotFound)
	}

	var t *transaction.Transaction

	err = i.repository.WithTransaction(ctx, func() error {
		payer.Debit(payload.Value)
		payee.Credit(payload.Value)

		if err := i.accountRepository.Save(ctx, payer); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}

		if err := i.accountRepository.Save(ctx, payee); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}

		t, err = transaction.New(payload.Value, payer, payee)
		if err != nil {
			return errors.New(http.StatusForbidden, err)
		}

		err = i.repository.Save(ctx, t)
		if err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}

		return nil
	})

	if err != nil {
		return mappers.TransactionDTO{}, err
	}

	return mappers.TransactionToDTO(t), nil
}
