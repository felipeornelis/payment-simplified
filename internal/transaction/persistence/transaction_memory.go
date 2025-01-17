package persistence

import (
	"context"
	"sync"

	"github.com/felipeornelis/payment-simplified/internal/transaction"
)

type memoryTransactionRepository struct {
	mutex        sync.RWMutex
	transactions map[string]*transaction.Transaction
}

func NewTransactionRepository() transaction.Repository {
	return &memoryTransactionRepository{
		transactions: make(map[string]*transaction.Transaction),
	}
}

// Save implements account.Repository.
func (r *memoryTransactionRepository) Save(ctx context.Context, tx *transaction.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.transactions[tx.ID()] = tx
	return nil
}

// WithTransaction implements transaction.Repository.
func (r *memoryTransactionRepository) WithTransaction(ctx context.Context, fn func() error) error {
	// Create a deep copy of the current state
	backup := make(map[string]*transaction.Transaction)
	for id, account := range r.transactions {
		prepare := *account
		backup[id] = &prepare
	}

	// Run the transaction
	if err := fn(); err != nil {
		// If an error occurs, revert to the backup
		r.transactions = backup
		return err
	}

	// Commit is automatic if no errors
	return nil
}
