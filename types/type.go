package types

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UpdatedResponseWriter struct {
	gin.ResponseWriter
	writes []string
	status int
}

func (ur *UpdatedResponseWriter) WriteString(s string) (int, error) {
	fmt.Println("WrittenString", s)
	ur.writes = append(ur.writes, s)
	return len(s), nil
}

func (ur *UpdatedResponseWriter) Status() int {
	return ur.status
}

func (ur *UpdatedResponseWriter) Flush() {
	for _, w := range ur.writes {
		_, err := ur.ResponseWriter.WriteString(w)
		fmt.Println(ur.writes)
		if err != nil {
			ur.ResponseWriter.Flush()
			return
		}
	}
	ur.ResponseWriter.Flush()
	return
}
