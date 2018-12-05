package middleware

import (
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("studentId", "123456")
		var data interface{}

		token := c.Query("token")
		if token == "" {
			logrus.Info("empty token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": e.INVALID_PARAMS,
				"msg":  e.GetMsg(e.INVALID_PARAMS),
				"data": data,
			})
			c.Abort()
			return
		} else {
			//从用户中心拉取用户信息，并设置到
			apiStudent, err := util.GetStudentFromUserCenter(token)
			if err != nil {
				logrus.WithField("token", token).Error(err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": e.ERROR_TOKEN_EXPIRE,
					"msg":  e.GetMsg(e.ERROR_TOKEN_EXPIRE),
					"data": data,
				})
				c.Abort()
				return
			}
			c.Set("studentId", apiStudent.Data.StudentId)
		}
		c.Next()
	}
}
