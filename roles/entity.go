package roles

import "pustaka-api/user"

type Roles struct {
	ID   uint
	Name string
	User user.User
}
