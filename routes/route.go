package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Home Page",
	})
}

func randomPanic() {
	panic("Crashed Application")
}

func PanicPage(ctx *gin.Context) {
	panic("Crashed Application")
}

func PanicAfterMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"Message": "Before a Panic",
	})
	panic("This is the Panic")
}

func PanicCase(ctx *gin.Context) {
	defer func() {
		err := recover()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err,
		})
	}()
	randomPanic()
}
