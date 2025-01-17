package transaction

import "context"

type Repository interface {
	Save(context.Context, *Transaction) error
	WithTransaction(context.Context, func() error) error
}
