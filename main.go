package main

import (
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
	dbName := os.Getenv("GB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)
	_, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error." + err.Error())
	}

	log.Println("Database Connection succeed")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	err = router.Run(":5000")
	helper.FatalIfError(err)
}
