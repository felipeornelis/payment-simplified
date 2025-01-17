package persistence

import (
	"context"
	"sync"

	"github.com/felipeornelis/payment-simplified/internal/account"
)

type memoryAccountRepository struct {
	mutex    sync.RWMutex
	accounts map[string]*account.Account
}

func NewMemoryAccountRepository() account.Repository {
	return &memoryAccountRepository{
		accounts: make(map[string]*account.Account),
	}
}

// FindByDocument implements account.Repository.
func (r *memoryAccountRepository) FindByDocument(ctx context.Context, document string) (*account.Account, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, acc := range r.accounts {
		if acc.Document() == document {
			return acc, nil
		}
	}

	return nil, account.ErrNotFound
}

// FindByEmail implements account.Repository.
func (r *memoryAccountRepository) FindByEmail(ctx context.Context, email string) (*account.Account, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, acc := range r.accounts {
		if acc.Email() == email {
			return acc, nil
		}
	}

	return nil, account.ErrNotFound
}

// FindByID implements account.Repository.
func (r *memoryAccountRepository) FindByID(ctx context.Context, id string) (*account.Account, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if acc, ok := r.accounts[id]; ok {
		return acc, nil
	}

	return nil, account.ErrNotFound
}

// Save implements account.Repository.
func (r *memoryAccountRepository) Save(ctx context.Context, acc *account.Account) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.accounts[acc.ID()] = acc
	return nil
}

// Update implements account.Repository.
func (r *memoryAccountRepository) Update(ctx context.Context, acc *account.Account) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.accounts[acc.ID()] = acc
	return nil
}

// Update implements account.Repository.
func (r *memoryAccountRepository) FindAll(ctx context.Context) []*account.Account {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var accounts []*account.Account

	for _, acc := range r.accounts {
		accounts = append(accounts, acc)
	}

	return accounts
}
