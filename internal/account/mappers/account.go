package mappers

import "github.com/felipeornelis/payment-simplified/internal/account"

type AccountDTO struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Type     account.Type `json:"type"`
	Document string       `json:"document"`
	Email    string       `json:"email"`
	Balance  float64      `json:"balance"`
}

func AccountToDTO(acc *account.Account) AccountDTO {
	return AccountDTO{
		ID:       acc.ID(),
		Name:     acc.Name(),
		Type:     acc.Type(),
		Balance:  acc.Balance(),
		Document: acc.Document(),
		Email:    acc.Email(),
	}
}
