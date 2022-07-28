package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Rizki Syaban Aryanto",
			"bio":  "Software Engineer",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":    "Hello Wolrd",
			"subtitle": "Belajar Golang API",
		})
	})

	router.Run(":5000")
}
