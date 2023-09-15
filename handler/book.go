package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, ok := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	RolesId, _ := claims["roles"].(int)
	fmt.Println(claims)
	fmt.Println(RolesId)
	if !ok || RolesId != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Admin role required."})
		return
	}

	// Jika pengguna adalah admin, maka lanjutkan dengan membuat buku
	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.ConvertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookById(c *gin.Context) {

	ID, _ := strconv.Atoi(c.Param("id"))

	b, err := h.bookService.FindById(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	bookResponse := book.ConvertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) FindByTitleHandler(c *gin.Context) {
	title := c.Param("title") // Get the book name from the URL parameter

	// Call the service to find the book by name
	b, err := h.bookService.FindByTitle(title)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Book not found",
		})
		return
	}
	bookResponse := book.ConvertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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
	ID, _ := strconv.Atoi(c.Param("id"))

	// jwtClaims, _ := c.Get("jwtClaims")
	// claims, _ := jwtClaims.(jwt.MapClaims)
	// userID, _ := claims["sub"].(float64)

	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID",
		})
		return
	}
	var bookFound bool
	for _, book := range books {
		if int(book.ID) == ID {
			bookFound = true
			break
		}
	}

	if !bookFound {
		c.JSON(http.StatusForbidden, gin.H{
			"errors": "Failed to update the book",
		})
		return
	}
	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	role, ok := claims["roles"].(float64)
	fmt.Println(claims)
	fmt.Println(role)
	roleID := uint(role)
	if !ok || roleID != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Admin role required."})
		return
	}


	b, err := h.bookService.Update(ID, bookRequest)
	bookResponse := book.ConvertToBookResponse(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateByTitleBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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
	title := c.Param("title")

	// jwtClaims, _ := c.Get("jwtClaims")
	// claims, _ := jwtClaims.(jwt.MapClaims)
	// userID, _ := claims["sub"].(float64)
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Book not found",
		})
		return
	}

	var bookFound bool
	for _, book := range books {
		if book.Title == title {
			bookFound = true
			break
		}
	}

	if !bookFound {
		c.JSON(http.StatusForbidden, gin.H{
			"errors": "Failed to update the book",
		})
		return
	}
	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, ok := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	rolesID, _ := claims["roles"].(uint)
	fmt.Println(claims)
	fmt.Println(rolesID)
	if !ok || rolesID != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Admin role required."})
		return
	}
	b, err := h.bookService.UpdateByTitle(title, bookRequest)
	bookResponse := book.ConvertToBookResponse(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	// jwtClaims, _ := c.Get("jwtClaims")
	// claims, _ := jwtClaims.(jwt.MapClaims)
	// userID, _ := claims["sub"].(float64)

	// books, err := h.bookService.FindAll(uint(userID))

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"errors": "Invalid ID",
	// 	})
	// 	return
	// }

	// var bookFound bool

	// for _, book := range books {
	// 	if int(book.ID) == ID {
	// 		bookFound = true
	// 		break
	// 	}
	// }
	// if !bookFound {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"errors": "Books not Found",
	// 	})
	// 	return
	// }
	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, ok := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	rolesID, _ := claims["roles"].(uint)
	fmt.Println(claims)
	fmt.Println(rolesID)
	if !ok || rolesID != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Admin role required."})
		return
	}
	b, err := h.bookService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "Failed to delete the book",
		})
		return
	}

	bookResponse := book.ConvertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

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

func (h *bookHandler) DeleteByTitleHandler(c *gin.Context) {
	title := c.Param("title")

	b, err := h.bookService.DeleteByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": "Failed to delete the book",
		})
		return
	}
	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, ok := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	rolesID, _ := claims["roles"].(uint)
	fmt.Println(claims)
	fmt.Println(rolesID)
	if !ok || rolesID != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied. Admin role required."})
		return
	}
	bookResponse := book.ConvertToBookResponse(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
