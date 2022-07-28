package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":5000")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Rizki Syaban Aryanto",
		"bio":  "Software Engineer",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Golang API",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"title":     bookInput.Title,
			"price":     bookInput.Price,
			"sub_title": bookInput.SubTitle,
		})
	}
}
