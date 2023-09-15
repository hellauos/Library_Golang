package main

import (
	// "encoding/json"

	// "net/http"

	"log"
	// "pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/account"
	"pustaka-api/initializer"

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
	// dsn := "root:@tcp(127.0.0.1:3306)/pustakaapi?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal("db connection error")
	// }

	// bookRepository := book.NewRepository(db)
	// bookService := book.NewService(bookRepository)
	// bookHandler := handler.NewBookHandler(bookService)

	accountRepository := account.NewRepository(db)
	accountService := account.NewService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	// // Creating Data===================
	// book := book.Book{}
	// book.Title = "American Captain"
	// book.Price = 57500
	// book.Description = "Hero of American People"
	// book.Rating = 9
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Creating Book Record")
	// }

	// Reading Data Up and Down Data===================
	// var book book.Book
	// // err = db.Debug().First(&book).Error
	// err = db.Debug().Last(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Finding Book Record")
	// }
	// fmt.Println("Book Object %v", book)

	// Reading Data Multiple===================
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }
	// for _, b := range books {
	// 	fmt.Println("book object %v", b)
	// }

	// Find with Condition===================
	// var books []book.Book
	// err = db.Debug().Where("title LIKE ?", "%ric%").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// }
	// for _, b := range books {
	// 	fmt.Println("book object %v", b)
	// }

	// Update Data - Full Update===================
	// var books book.Book
	// err = db.Debug().First(&books, 0).Error
	// if err != nil {
	// 	fmt.Println("Error finding book record")
	// } else {
	// 	books.Title = "Atomic Habitats"
	// 	err = db.Save(&books).Error
	// 	if err != nil {
	// 		fmt.Println("Error Update Book")
	// 	}
	// }

	// Delete Data - Full Update===================
	// var books book.Book
	// books.ID = 5
	// err = db.Delete(&books).Error

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

	routerV1.POST("/SignUp", accountHandler.SignUp)
	routerV1.GET("/Account/Role", accountHandler.GetAccountByRole)
	routerV1.POST("/Login/Auth", accountHandler.Login)

	// routerV1Books := routerV1.Group("/books", middleware.RequiredAuth)
	// routerV1Books.POST("", bookHandler.PostBooksHandler)
	// routerV1Books.GET("", bookHandler.GetBooksByUser)
	// routerV1Books.GET("/:id", bookHandler.GetBookById)
	// routerV1Books.PUT("/:id", bookHandler.UpdateBookHandler)
	// routerV1Books.DELETE("/:id", bookHandler.DeleteBookHandler)


	

	router.Run(":3030")
}
