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

	// Create
	createDataBook := book.Book{}
	createDataBook.Title = "Gundam Blood Orphan 2"
	createDataBook.Price = 20000
	createDataBook.Discount = 5
	createDataBook.Rating = 8
	createDataBook.Description = "Ini buku bagus banget banget gais"

	err = db.Create(&createDataBook).Error
	helper.FatalIfError(err)

	// Read
	var dataBook book.Book

	err = db.Debug().First(&dataBook, 2).Error
	helper.FatalIfError(err)

	log.Println("Title: " + dataBook.Title)

	var dataBooks []book.Book

	err = db.Debug().Where("title like ?", "%Gundam Blood Orphan%").Find(&dataBooks).Error
	helper.FatalIfError(err)

	for _, b := range dataBooks {
		log.Println("Title: " + b.Title)
	}

	// Update
	var dataUpdateBook book.Book

	err = db.Debug().Where("id = ?", 1).First(&dataUpdateBook).Error
	helper.FatalIfError(err)

	dataUpdateBook.Title = "Gundam Revised"
	err = db.Save(&dataUpdateBook).Error
	helper.FatalIfError(err)

	// Delete
	var deleteDataBook book.Book

	err = db.Debug().Where("id = ?", 1).First(&deleteDataBook).Error
	helper.FatalIfError(err)

	err = db.Delete(&deleteDataBook).Error
	helper.FatalIfError(err)

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
