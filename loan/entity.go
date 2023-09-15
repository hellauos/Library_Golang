package loan

import (
	"time"
)

type Loan struct {
	ID      uint
	DueDate time.Duration
	Status  int
	BookId  uint
	UserId  uint
}
