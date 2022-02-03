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
		defaultRoutes.GET("/paniclater", panicAfterMessage)
		defaultRoutes.GET("/panicrandom", panicCase)
	}
	router.Run(":5000")
}

func defaultPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Home Page",
	})
}

func randomPanic() {
	panic("Crashed Application")
}

func panicPage(ctx *gin.Context) {
	randomPanic()
}

func panicAfterMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Before a Panic",
	})
	panic("This is the Panic")
}

func panicCase(ctx *gin.Context) {
	defer func() {
		err := recover()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
	}()
	randomPanic()
}
