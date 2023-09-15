package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(signUpRequest SignUpRequest) (User, error)
	Login(loginRequest LoginRequest) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(signUpRequest SignUpRequest) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 10)
	if err != nil {
		return User{}, err
	}
	// Save user
	user := User{
		Name:     signUpRequest.Name,
		Email:    signUpRequest.Email,
		Password: string(hash),
		RolesId:  signUpRequest.RolesID,
	}
	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Login(loginRequest LoginRequest) (string, error) {
	//get user
	user, err := s.repository.FindByEmail(loginRequest.Email)

	if err != nil {
		return "", err
	} else if user.ID == 0 {
		return "", errors.New("Invalid Email or Password")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		return "", err
	}

	//sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"roles": user.RolesId,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
