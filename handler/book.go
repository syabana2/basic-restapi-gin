package handler

import (
	"basic-rest-api-gin/book"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Rizki Syaban Aryanto",
		"bio":  "Software Engineer",
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Golang API",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput book.Input
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
