package account

type SignupRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RolesId uint `json:"roles_id"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
