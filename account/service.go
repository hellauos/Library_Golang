package account

import (
	"errors"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(signupRequest SignupRequest) (Account, error)
	Login(loginRequest LoginRequest) (string, error)
	GetAccountByRole(roleID string) ([]Account, error)
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
		Name:     signupRequest.Name,
		Email:    signupRequest.Email,
		Password: string(hash),
		RolesId: signupRequest.RolesId,
	}

	newAccount, err := s.repository.Create(account)
	return newAccount, nil
}

func (s *service) Login(loginRequest LoginRequest) (string, error) {
	account, err := s.repository.GetAccountByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}else if account.ID == 0{
		return"", errors.New("Invalid email and password")
	}

	//Verifikasi Password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginRequest.Password))
	if err != nil {
		return "",err
	}

// 	//Membuat Token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_Id": account.ID,
		"roles": account.RolesId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	//Return JWT sebagai respons
	return tokenString, nil
}


func (s *service) GetAccountByRole(roleName string) ([]Account, error) {
    // Panggil metode repository yang sesuai untuk mengambil akun berdasarkan peran (role)
    account, err := s.repository.GetAccountByRole(roleName)
    if err != nil {
        return nil, err
    }
    
    return account, nil
}

// GenerateToken menghasilkan token JWT untuk pengguna yang berhasil login.
// func GenerateToken(accountID uint, secretKey string, roleID string) (string, error) {
// 	//Membuat claim token
// 	claims := jwt.MapClaims{
// 		"account_id": accountID,
// 		"role_id": roleID,
// 		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
// 	}
// 	// Membuat token dengan klaim
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Menandatangani token dengan kunci rahasia
// 	tokenString, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil

// }
