package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Quons/go-gin-example/pkg/e"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			c.Set("studentId",int64(123456))
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
