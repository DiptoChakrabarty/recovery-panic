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
		defaultRoutes.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Home Page",
			})
		})
	}
	router.Run(":5000")
}
