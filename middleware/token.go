package middleware

import (
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/Quons/go-gin-example/models"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("studentId", "123456")
		var data interface{}

		token := c.Query("token")
		logrus.Info("token:",token)
		if token == "" {
			logrus.Info("empty token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": e.ERROR_INVALID_PARAMS,
				"msg":  e.GetMsg(e.ERROR_INVALID_PARAMS),
				"data": data,
			})
			c.Abort()
			return
		} else {

			//从用户中心拉取用户信息，并设置到
			apiStudent, err := util.GetStudentFromUserCenter(token)
			logrus.Info("apiStudent:","......")
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
			studentInfo, err := models.GetStudent(apiStudent.Data.StudentId)
			if err != nil || studentInfo.StudentID == 0 {
				logrus.WithField("token", token).Error(err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": e.ERROR_TOKEN_EXPIRE,
					"msg":  e.GetMsg(e.ERROR_TOKEN_EXPIRE),
					"data": data,
				})
				c.Abort()
				return
			}
			c.Set(gin.AuthUserKey, studentInfo)
		}
		c.Next()
	}
}
