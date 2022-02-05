package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/DiptoChakrabarty/recovery-panic/types"
	"github.com/gin-gonic/gin"
)

func RecoverPanic(dev bool) gin.HandlerFunc {
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
		nrw := &types.UpdatedResponseWriter{ResponseWriter: ctx.Writer}
		nrw.Flush()
		ctx.Writer = nrw
		ctx.Next()
	}
}
