package transaction

import (
	"time"

	"github.com/felipeornelis/payment-simplified/internal/account"
	"github.com/google/uuid"
)

type Transaction struct {
	id        string
	value     float64
	payer     *account.Account
	payee     *account.Account
	createdAt time.Time
}

func New(value float64, payer, payee *account.Account) (*Transaction, error) {
	transaction := &Transaction{
		id:        uuid.New().String(),
		value:     value,
		payer:     payer,
		payee:     payee,
		createdAt: time.Now(),
	}

	if err := transaction.validate(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) validate() error {
	if t.payer.Balance() < t.value {
		return ErrInsufficientBalance
	}

	if t.payer.Type() == account.SellerType {
		return ErrSellerPayerInvalidOperation
	}

	return nil
}

func (t *Transaction) ID() string              { return t.id }
func (t *Transaction) Value() float64          { return t.value }
func (t *Transaction) Payer() *account.Account { return t.payer }
func (t *Transaction) Payee() *account.Account { return t.payee }
func (t *Transaction) CreatedAt() time.Time    { return t.createdAt }
