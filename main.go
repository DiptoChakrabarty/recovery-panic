package main

import (
	"github.com/DiptoChakrabarty/recovery-panic/middleware"
	"github.com/DiptoChakrabarty/recovery-panic/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	var dev bool
	dev = true
	defaultRoutes := router.Group("/", middleware.RecoverPanic(dev))
	{
		defaultRoutes.GET("", routes.DefaultPage)
		defaultRoutes.GET("/panic", routes.PanicPage)
		defaultRoutes.GET("/paniclater", routes.PanicAfterMessage)
		defaultRoutes.GET("/panicrandom", routes.PanicCase)
	}
	router.Run(":5000")
}
