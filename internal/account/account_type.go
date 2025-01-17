package account

import "encoding/json"

type Type uint

const (
	UserType = iota + 1
	SellerType
)

func (t Type) String() string {
	switch t {
	case UserType:
		return "user"
	case SellerType:
		return "seller"
	default:
		return "unknown_account_type"
	}
}

func (t Type) Valid() bool {
	switch t {
	case UserType, SellerType:
		return true
	}

	return false
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var typeMap = map[string]Type{
	"user":   UserType,
	"seller": SellerType,
}

func ParseType(kind string) (Type, error) {
	value, ok := typeMap[kind]
	if ok {
		return value, nil
	}

	return 0, ErrUnknownAccountType
}
