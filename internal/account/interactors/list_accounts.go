package interactors

import (
	"context"

	"github.com/felipeornelis/payment-simplified/internal/account"
	"github.com/felipeornelis/payment-simplified/internal/account/mappers"
)

type ListAccountsInteractor interface {
	Execute(context.Context) []mappers.AccountDTO
}

type listAccountsInteractor struct {
	repository account.Repository
}

func NewListAccountsInteractor(r account.Repository) ListAccountsInteractor {
	return &listAccountsInteractor{
		repository: r,
	}
}

// Execute implements CreateAccountInteractor.
func (i *listAccountsInteractor) Execute(ctx context.Context) []mappers.AccountDTO {
	var dtos []mappers.AccountDTO

	for _, acc := range i.repository.FindAll(ctx) {
		dtos = append(dtos, mappers.AccountToDTO(acc))
	}

	return dtos
}
