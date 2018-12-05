package v1

import (
	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/service/course_service"
	"github.com/Quons/go-gin-example/vo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Tags 课程
// @Summary 获取单个课程
// @Description 获取单个课程description
// @Produce  json
// @Param token query string true "用户token"
// @Param courseId query int true "课程ID"
// @Success 200 {object} vo.CourseVo
// @Failure 10000 {string} json "{"code":10000,"data":{},"msg":"服务器错误"}"
// @Failure 20000 {string} json "{"code":20000,"data":{},"msg":"参数错误"}"
// @Router /api/v1/getCourse [post]
func GetCourse(c *gin.Context) {
	appG := app.Gin{C: c}
	course := course_service.Course{}
	err := c.Bind(&course)
	if err != nil {
		logrus.Info(err)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}
	studentId := c.GetInt64("studentId")
	logrus.WithField("studentId", studentId)
	courseDO, err := course.Get()
	if err != nil {
		appG.Response(nil, e.ERROR_SERVER_ERROR)
		return
	}
	if courseDO.CourseID == 0 {
		appG.Response(nil, e.ERROR_DATA_ERROR)
		return
	}
	appG.Response(vo.CourseVo{}.Transform(courseDO), e.SUCCESS)
}
