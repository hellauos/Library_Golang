package roles

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Name      string
	AccountId uint
}
