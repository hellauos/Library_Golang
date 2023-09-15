package loan

type LoanBookRequest struct {
	AccountId int    `json:"account_id" binding:"required"`
	BookId    int    `json:"book_id" binding:"required"`
	DueDate   string `json:"due_date" binding:"required"`
}

type ReturnBookRequest struct {
	LoanId int `json:"loan_id" binding:"required"`
}

type Book struct {
	ID          uint
	Title       string
	Description string
	Location    string
	Price       int
	Stock       int
	CategoryID  uint
}

type Account struct {
	ID       uint
	name     string
	Email    string `gorm:"unique"`
	Password string
	RolesId  uint
}
