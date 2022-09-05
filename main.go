package main

import (
	"log"
	"pustakaapi/book"
	"pustakaapi/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustakaapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	// migrate table from model
	// db.AutoMigrate(&book.Book{})

	bookRepository := book.NeRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRequest := book.BookRequest{
	// 	Title:       "tendang gundam",
	// 	Price:       "9090909",
	// 	Description: "dshjdsd",
	// 	Discount:    23,
	// 	Rating:      3,
	// }
	// bookService.Create(bookRequest)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	v1.POST("/books", bookHandler.PostBookHandler)
	router.Run(":8080")
}
