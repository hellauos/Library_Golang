package account

type SignupRequest struct {
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}