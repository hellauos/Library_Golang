package main

import (
	"log"
	"pustaka-api/account"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initializer"
	"pustaka-api/loan"
	"pustaka-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	accountRepository := account.NewRepository(db)
	accountService := account.NewService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

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
	routerV1.POST("/signUp", accountHandler.SignUp)

	routerV1.POST("/login", accountHandler.Login)

	routerV1Books := routerV1.Group("/books", middleware.RequiredAuth)
	routerV1Books.POST("", bookHandler.PostBooksHandler)
	routerV1Books.GET("", bookHandler.GetBooks)
	routerV1Books.GET("/id/:id", bookHandler.GetBookById)
	routerV1Books.GET("/title/:title", bookHandler.FindByTitleHandler)
	routerV1Books.PUT("/id/:id", bookHandler.UpdateBookHandler)
	routerV1Books.PUT("/title/:title", bookHandler.UpdateByTitleBookHandler)
	routerV1Books.DELETE("/id/:id", bookHandler.DeleteBookHandler)
	routerV1Books.DELETE("/title/:title", bookHandler.DeleteByTitleHandler)

	router.Run(":3030")
}
