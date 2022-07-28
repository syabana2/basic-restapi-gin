package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

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
