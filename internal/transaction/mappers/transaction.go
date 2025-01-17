package mappers

import (
	"time"

	"github.com/felipeornelis/payment-simplified/internal/transaction"
)

type TransactionDTO struct {
	ID        string    `json:"id"`
	Value     float64   `json:"value"`
	PayerID   string    `json:"payer"`
	PayeeID   string    `json:"payee"`
	CreatedAt time.Time `json:"created_at"`
}

func TransactionToDTO(t *transaction.Transaction) TransactionDTO {
	return TransactionDTO{
		ID:        t.ID(),
		Value:     t.Value(),
		PayerID:   t.Payer().ID(),
		PayeeID:   t.Payee().ID(),
		CreatedAt: t.CreatedAt(),
	}
}
