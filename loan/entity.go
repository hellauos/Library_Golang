package loan

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	DueDate time.Duration
	Status  int
	Book    int
	Account int
}
