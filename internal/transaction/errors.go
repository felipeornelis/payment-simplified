package transaction

import "errors"

var (
	ErrInsufficientBalance         = errors.New("insufficient balance to complete the transaction")
	ErrSellerPayerInvalidOperation = errors.New("seller cannot transfer money")
)
