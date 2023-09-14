package loan

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	DueDate   time.Duration
	Status    int
	BookId    uint
	AccountId uint
}
