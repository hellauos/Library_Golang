package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initializer"
	"pustaka-api/middleware"
	"pustaka-api/account"

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

	accountRepository := account.NewRepository(db)
	accountService := account.NewService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)
	router := gin.Default()

	routerV1 := router.Group("/v1")

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
