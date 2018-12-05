package v1

import (
	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/service/course_service"
	"github.com/Quons/go-gin-example/vo"
	"github.com/Unknwon/com"
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
// @Router /api/v2/getCourse [post]
func GetCourse(c *gin.Context) {
	appG := app.Gin{C: c}
	courseID, err := com.StrTo(c.Query("courseId")).Int64()
	if err != nil {
		logrus.Info(err)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}

	articleService := &course_service.Course{CourseId: courseID}
	course, err := articleService.Get()
	if err != nil {
		appG.Response(nil, e.ERROR_SERVER_ERROR)
		return
	}
	if course.CourseID == 0 {
		appG.Response(nil, e.ERROR_DATA_ERROR)
		return
	}
	appG.Response(vo.CourseVo{}.Transform(course), e.SUCCESS)
}
