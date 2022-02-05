package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/DiptoChakrabarty/recovery-panic/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	var dev bool
	dev = true
	defaultRoutes := router.Group("/", recoverPanic(dev))
	{
		defaultRoutes.GET("", routes.DefaultPage)
		defaultRoutes.GET("/panic", routes.PanicPage)
		defaultRoutes.GET("/paniclater", routes.PanicAfterMessage)
		defaultRoutes.GET("/panicrandom", routes.PanicCase)
	}
	router.Run(":5000")
}

func recoverPanic(dev bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"Message": "Error Occured",
					})
				} else {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"Message": err,
					})
				}
			}
		}()
		ctx.Next()
	}
}
