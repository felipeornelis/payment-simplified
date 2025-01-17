package account

import "errors"

var (
	ErrUnknownDocumentType    = errors.New("unknown document type")
	ErrUnknownAccountType     = errors.New("unknown account type")
	ErrNotFound               = errors.New("account not found")
	ErrEmailAddressRegistered = errors.New("email address is registered")
	ErrCNPJRegistered         = errors.New("CNPJ number is registered")
	ErrCPFRegistered          = errors.New("CPF number is registered")
)
