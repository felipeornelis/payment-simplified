package account

import "context"

type Repository interface {
	Save(context.Context, *Account) error
	Update(context.Context, *Account) error
	FindByDocument(context.Context, string) (*Account, error)
	FindByEmail(context.Context, string) (*Account, error)
	FindByID(context.Context, string) (*Account, error)
	FindAll(context.Context) []*Account
}
