package handler

import (
	"basic-rest-api-gin/book"
	"basic-rest-api-gin/helper"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Rizki Syaban Aryanto",
		"bio":  "Software Engineer",
	})
}

func (h *bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Golang API",
	})
}

func (h *bookHandler) BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) QueryHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) PostBooksHandler(ctx *gin.Context) {
	var bookRequest book.Request
	var price int64

	err := ctx.ShouldBindJSON(&bookRequest)
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

	_, status := bookRequest.Price.(string)
	if status {
		price, err = strconv.ParseInt(bookRequest.Price.(string), 10, 64)
		helper.FatalIfError(err)
	} else {
		price = int64(bookRequest.Price.(float64))
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(price)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "success",
		"data":   book,
	})
}
