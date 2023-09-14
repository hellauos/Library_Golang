package account

import (
	"errors"
	"time"
	"strconv"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(signupRequest SignupRequest) (Account, error)
	Login(loginRequest LoginRequest) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(signupRequest SignupRequest) (Account, error) {

	//Hash the password
	hash, err :=
		bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 10)

	if err != nil {
		return Account{}, err
	}

	//Save account
	account := Account{
		name:     signupRequest.Name,
		Email:    signupRequest.Email,
		Password: string(hash),
	}

	newAccount, err := s.repository.Create(account)
	return newAccount, err
}

func (s *service) Login(loginRequest LoginRequest) (LoginResponse, error) {
	account, err := s.repository.GetAccountByEmail(loginRequest.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	//Verifikasi Password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginRequest.Password))
	if err != nil {
		return LoginResponse{}, errors.New("Invalid credentials")
	}

	roleID := account.RolesId

	//Membuat Token JWT
	token, err := GenerateToken(account.ID, "asdasd", strconv.Itoa(int(roleID)))
	if err != nil {
		return LoginResponse{}, err
	}

	//Return JWT sebagai respons
	return LoginResponse{
		Token: token,
	}, nil
}

// GenerateToken menghasilkan token JWT untuk pengguna yang berhasil login.
func GenerateToken(accountID uint, secretKey string, roleID string) (string, error) {
	//Membuat claim token
	claims := jwt.MapClaims{
		"user_id": accountID,
		"role_id": roleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	}
	// Membuat token dengan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
