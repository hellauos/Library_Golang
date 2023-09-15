package handler

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	// "pustaka-api/book"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	genre := c.Query("genre")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"genre": genre,
	})
}
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "With RootHandler",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
	})
}

// // func PostBooksHandler(c *gin.Context) {
// // 	var bookInput book.BookRequest

// // 	err := c.ShouldBindJSON(&bookInput)
// // 	if err != nil {
// // 		switch err.(type) {
// // 		case validator.ValidationErrors:
// // 			errorMessages := []string{}
// // 			for _, e := range err.(validator.ValidationErrors) {
// // 				errorMessage := fmt.Sprintf("Error on Field: %s, condition: %s", e.Field(), e.ActualTag())
// // 				errorMessages = append(errorMessages, errorMessage)
// // 			}

// // 			c.JSON(http.StatusBadRequest, gin.H{
// // 				"errors": errorMessages,
// // 			})
// // 			return
// // 		case *json.UnmarshalTypeError:
// // 			c.JSON(http.StatusBadRequest, gin.H{
// // 				"errors": err.Error(),
// // 			})
// // 			return
// // 		}

// // 		c.JSON(http.StatusBadRequest, err)
// // 		fmt.Println(err)
// // 		return

// // 		log.Fatal(err)
// // 	}

// // 	c.JSON(http.StatusOK, gin.H{
// // 		"title": bookInput.Title,
// // 		"price": bookInput.Price,
// // 	})
// }
