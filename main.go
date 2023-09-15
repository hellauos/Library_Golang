package main

import (
	// "encoding/json"

	// "net/http"

	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initializer"
	"pustaka-api/loan"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
)

var db *gorm.DB

func init() {
	var err error

	initializer.LoadEnvVariables()
	db, err = initializer.ConnectToDatabase()
	err = initializer.SyncDatabase(db)

	if err != nil {
		log.Fatal("db connection error/ failed")
	}
}

func main() {

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	loanRepository := loan.NewRepository(db)
	loanService := loan.NewService(loanRepository)
	loanHandler := handler.NewLoanHandler(loanService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	// routerV1.GET("/", handler.RootHandler)

	// routerV1.GET("/hello", handler.HelloHandler)

	// router.GET("/books/:id", booksHandler)

	// routerV1.GET("/books/:id/:title", handler.BooksHandler)

	// routerV1.GET("/query", handler.QueryHandler)

	// routerV1.POST("/books", bookHandler.PostBooksHandler)
	// routerV1.GET("/books", bookHandler.GetBooks)
	// routerV1.GET("/books/:id", bookHandler.GetBookById)
	// routerV1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	// routerV1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	// routerV1.POST("/signUp", userHandler.SignUp)

	// routerV1.POST("/login", userHandler.Login)

	// routerV1Books := routerV1.Group("/books", middleware.RequiredAuth)
	routerV1.POST("/books/GetBookByTitleCategory", bookHandler.GetBookByTitleCategory)
	routerV1.POST("/books/LoanBook", loanHandler.LoanBook)
	routerV1.POST("/books/GetLoanData", loanHandler.GetLoanData)
	routerV1.POST("/books/ReturnBook", loanHandler.ReturnBook)
	
	// routerV1Books.POST("", bookHandler.PostBooksHandler)
	// routerV1Books.GET("/:id", bookHandler.GetBookById)
	// routerV1Books.PUT("/:id", bookHandler.UpdateBookHandler)
	// routerV1Books.DELETE("/:id", bookHandler.DeleteBookHandler)

	router.Run(":3030")
}
