package main

import (
	"basic-rest-api-gin/book"
	"basic-rest-api-gin/handler"
	"basic-rest-api-gin/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	helper.FatalIfError(err)

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error." + err.Error())
	}
	err = db.AutoMigrate(&book.Book{})
	helper.FatalIfError(err)

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	err = router.Run(":5000")
	helper.FatalIfError(err)
}
