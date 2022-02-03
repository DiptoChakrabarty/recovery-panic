package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	defaultRoutes := router.Group("/")
	{
		defaultRoutes.GET("", defaultPage)
		defaultRoutes.GET("/panic", panicPage)
	}
	router.Run(":5000")
}

func defaultPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Home Page",
	})
}

func panicPage(ctx *gin.Context) {
	panic("Crashed Application")
}
