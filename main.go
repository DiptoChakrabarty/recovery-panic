package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	defaultRoutes := router.Group("/", recoverPanic())
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

func recoverPanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Error Occured",
				})
			}
		}()
		ctx.Next()
	}
}

func randomPanic() {
	panic("Crashed Application")
}

func panicPage(ctx *gin.Context) {
	panic("Crashed Application")
}

func panicAfterMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
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
