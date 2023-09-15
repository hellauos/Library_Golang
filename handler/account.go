package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/account"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type accountHandler struct {
	bookService account.Service
}

func NewAccountHandler(service account.Service) *accountHandler {
	return &accountHandler{service}
}

func (h *accountHandler) SignUp(c *gin.Context) {
	var signUpRequest account.SignupRequest

	err := c.ShouldBindJSON(&signUpRequest)
	// tanda & menunjukkan itu adalah pointernya, sedangakan tanda * merupakan valuenya

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}
	fmt.Println(signUpRequest)
	account, err := h.bookService.Create(signUpRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Account %v created successfully", account.Email),
	})
}

func (h *accountHandler) Login(c *gin.Context) {
	var loginRequest account.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	tokenString, err := h.bookService.Login(loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"token": tokenString,
		},
	})
}

func (h *accountHandler) GetAccountByRoleHandler(c *gin.Context) {
    roleID := c.Param("roleID") // Ambil roleID dari URL atau request parameter sesuai dengan framework yang Anda gunakan
    
	account, err := h.service.GetAccountByRole(roleID)
    if err != nil {
        // Handle kesalahan, misalnya dengan mengirimkan respons HTTP 500
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
        return
    }

	// Kirim respons HTTP 200 dengan daftar akun yang sesuai
    c.JSON(http.StatusOK, account)
}
