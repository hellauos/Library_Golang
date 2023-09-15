package loan

import (
	"time"
)

type Loan struct {
	ID        uint
	DueDate   time.Time
	Status    int
	BookId    uint
	AccountId uint
}
