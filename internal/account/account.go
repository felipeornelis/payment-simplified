package account

import "github.com/google/uuid"

type Account struct {
	id       string
	name     string
	document string
	email    string
	kind     Type
	password string
	balance  float64
}

func New(name, email, password, document string, balance float64, kind Type) (*Account, error) {
	account := &Account{
		id:       uuid.New().String(),
		name:     name,
		document: document,
		email:    email,
		kind:     kind,
		password: password,
		balance:  balance,
	}

	if err := account.validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) validate() error {
	if len(a.document) != CPFDigits && len(a.document) != CNPJDigits {
		return ErrUnknownDocumentType
	}

	return nil
}

func (a *Account) Debit(value float64) {
	a.balance -= value
}

func (a *Account) Credit(value float64) {
	a.balance += value
}

func (a *Account) ID() string       { return a.id }
func (a *Account) Name() string     { return a.name }
func (a *Account) Document() string { return a.document }
func (a *Account) Email() string    { return a.email }
func (a *Account) Type() Type       { return a.kind }
func (a *Account) Password() string { return a.password }
func (a *Account) Balance() float64 { return a.balance }
