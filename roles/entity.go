package roles

import "pustaka-api/account"

type Roles struct {
	ID   uint
	Name string
	Account account.Account
}
