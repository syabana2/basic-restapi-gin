package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBooksHandler)

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
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
	//SubTitle string      `json:"sub_title"`
}

func postBooksHandler(ctx *gin.Context) {
	var bookInput BookInput
	var price int64

	err := ctx.BindJSON(&bookInput)
	if err != nil {
		_, status := err.(*json.SyntaxError)
		if status {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			var errorMessages []string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessages,
			})
			return
		}
	}

	_, status := bookInput.Price.(string)
	if status {
		price, err = strconv.ParseInt(bookInput.Price.(string), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		price = int64(bookInput.Price.(float64))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": price,
		//"sub_title": bookInput.SubTitle,
	})
}
