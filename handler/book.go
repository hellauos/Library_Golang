package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) GetBookByTitleCategory(c *gin.Context) {
	var getBookByTitleCategoryRequest book.GetBookByTitleCategoryRequest
	err := c.ShouldBindJSON(&getBookByTitleCategoryRequest)
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
	book, err := h.bookService.FindBookByTitleCategory(getBookByTitleCategoryRequest)
	if err != nil {
		errorMessage := err.Error() // Extract the error message
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// func (h *bookHandler) PostBooksHandler(c *gin.Context) {
// 	var bookRequest book.BookRequest

// 	err := c.ShouldBindJSON(&bookRequest)

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Filed: %s, condition: %s", e.Field(), e.ActualTag())
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

// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)
// 	book, err := h.bookService.Create(bookRequest, uint(userID))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": book,
// 	})
// }

// func (h *bookHandler) GetBooks(c *gin.Context) {
// 	books, err := h.bookService.FindAll()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}

// 	var booksResponse []book.BookResponse

// 	for _, b := range books {
// 		bookResponse := book.ConvertToBookResponse(b)
// 		booksResponse = append(booksResponse, bookResponse)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": booksResponse,
// 	})
// }

// func (h *bookHandler) GetBookById(c *gin.Context) {

// 	ID, _ := strconv.Atoi(c.Param("id"))

// 	b, err := h.bookService.FindById(ID)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}

// 	bookResponse := book.ConvertToBookResponse(b)

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": bookResponse,
// 	})
// }

// func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
// 	var bookRequest book.BookRequest

// 	err := c.ShouldBindJSON(&bookRequest)

// 	if err != nil {
// 		switch err.(type) {
// 		case validator.ValidationErrors:
// 			errorMessages := []string{}
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errorMessage := fmt.Sprintf("Error on Filed: %s, condition: %s", e.Field(), e.ActualTag())
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
// 	ID, _ := strconv.Atoi(c.Param("id"))

// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)

// 	books, err := h.bookService.FindAllBooksByUser(uint(userID))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": "Invalid ID",
// 		})
// 		return
// 	}
// 	var bookFound bool
// 	for _, book := range books {
// 		if int(book.Id) == ID {
// 			bookFound = true
// 			break
// 		}
// 	}

// 	if !bookFound {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"errors": "Books not Found",
// 		})
// 		return
// 	}

// 	b, err := h.bookService.Update(ID, bookRequest)
// 	bookResponse := book.ConvertToBookResponse(b)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": bookResponse,
// 	})
// }

// func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
// 	ID, err := strconv.Atoi(c.Param("id"))

// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)

// 	books, err := h.bookService.FindAllBooksByUser(uint(userID))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": "Invalid ID",
// 		})
// 		return
// 	}

// 	var bookFound bool

// 	for _, book := range books {
// 		if int(book.Id) == ID {
// 			bookFound = true
// 			break
// 		}
// 	}
// 	if !bookFound {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"errors": "Books not Found",
// 		})
// 		return
// 	}

// 	b, err := h.bookService.Delete(ID)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"errors": "Failed to delete the book",
// 		})
// 		return
// 	}

// 	bookResponse := book.ConvertToBookResponse(b)
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": bookResponse,
// 	})
// }

// func (h *bookHandler) GetBooksByUser(c *gin.Context) {
// 	jwtClaims, _ := c.Get("jwtClaims")
// 	claims, _ := jwtClaims.(jwt.MapClaims)
// 	userID, _ := claims["sub"].(float64)

// 	books, err := h.bookService.FindAllBooksByUser(uint(userID))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}

// 	var booksResponse []book.BookResponse

// 	for _, b := range books {
// 		bookResponse := book.ConvertToBookResponse(b)
// 		booksResponse = append(booksResponse, bookResponse)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": booksResponse,
// 	})
// }
