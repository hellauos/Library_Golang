package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/loan"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type loanHandler struct {
	loanService loan.Service
}

func NewLoanHandler(service loan.Service) *loanHandler {
	return &loanHandler{service}
}

func (h *loanHandler) LoanBook(c *gin.Context) {
	var loanRequest loan.LoanBookRequest
	err := c.ShouldBindJSON(&loanRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Filed: %s, condition: %s", e.Field(), e.ActualTag())
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

	loan, err := h.loanService.LoanBook(loanRequest)
	if err != nil {
		errorMessage := err.Error() // Extract the error message
		fmt.Println("handler", errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": loan,
	})
}

func (h *loanHandler) GetLoanData(c *gin.Context) {
	loan, err := h.loanService.FindLoanData()
	if err != nil {
		errorMessage := err.Error() // Extract the error message
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": loan,
	})
}

func (h *loanHandler) ReturnBook(c *gin.Context) {
	var returnBookRequest loan.ReturnBookRequest
	//validate
	err := c.ShouldBindJSON(&returnBookRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on Filed: %s, condition: %s", e.Field(), e.ActualTag())
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

	//main logic
	loan, err := h.loanService.ReturnBook(returnBookRequest)
	if err != nil {
		errorMessage := err.Error() // Extract the error message
		fmt.Println("handler", errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": loan,
	})
}
