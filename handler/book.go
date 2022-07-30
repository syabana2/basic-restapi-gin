package handler

import (
	"basic-rest-api-gin/book"
	"basic-rest-api-gin/helper"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	var booksResponse []book.Response

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	dataBook, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	bookResponse := convertToBookResponse(dataBook)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.Request

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		_, status := err.(*json.SyntaxError)
		if status {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			var errorMessages []string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errorMessages,
			})
			return
		}
	}

	price := helper.ConvertInterfaceToInt(bookRequest.Price)
	rating := helper.ConvertInterfaceToInt(bookRequest.Rating)
	discount := helper.ConvertInterfaceToInt(bookRequest.Discount)

	dataBook, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var bookResponse book.Response

	bookResponse = book.Response{
		ID:          dataBook.ID,
		Title:       dataBook.Title,
		Description: dataBook.Description,
		Price:       price,
		Rating:      rating,
		Discount:    discount,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "success",
		"data":   bookResponse,
	})
}

func convertToBookResponse(dataBook book.Book) book.Response {
	return book.Response{
		ID:          dataBook.ID,
		Title:       dataBook.Title,
		Description: dataBook.Description,
		Price:       dataBook.Price,
		Rating:      dataBook.Rating,
		Discount:    dataBook.Discount,
	}
}
