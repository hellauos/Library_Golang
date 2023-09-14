package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	user "pustaka-api/account"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// )

// type userHandler struct {
// 	bookService user.Service
// }

// func NewUserHandler(service user.Service) *userHandler {
// 	return &userHandler{service}
// }

// func (h *userHandler) SignUp(c *gin.Context) {
// 	var signUpRequest user.SignUpRequest

// 	err := c.ShouldBindJSON(&signUpRequest)
// 	// tanda & menunjukkan itu adalah pointernya, sedangakan tanda * merupakan valuenya

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
// 				errorMessages = append(errorMessages, errorMessage)
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": errorMessages,
// 			})
// 			return
// 		case *json.UnmarshalTypeError:
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": err.Error(),
// 			})
// 			return
// 		}
// 	}
// 	fmt.Println(signUpRequest)
// 	user, err := h.bookService.Create(signUpRequest)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("User %v created successfully", user.Email),
// 	})
// }

// func (h *userHandler) Login(c *gin.Context) {
// 	var loginRequest user.LoginRequest

// 	err := c.ShouldBindJSON(&loginRequest)

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
// 				errorMessages = append(errorMessages, errorMessage)
// 			}
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": errorMessages,
// 			})
// 			return
// 		case *json.UnmarshalTypeError:
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"errors": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	tokenString, err := h.bookService.Login(loginRequest)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": map[string]string{
// 			"token": tokenString,
// 		},
// 	})
// }
