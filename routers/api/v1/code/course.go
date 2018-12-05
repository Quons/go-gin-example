package code

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/service/course_service"
	"github.com/Quons/go-gin-example/vo"
	"github.com/sirupsen/logrus"
)

// @Summary 获取单个课程
// @Produce  json
// @Param token query string true "用户token"
// @Param courseId query int true "课程ID"
// @Success 200 {string} json "{"code":200,"data":{"id":3,"created_on":1516937037,"modified_on":0,"tag_id":11,"tag":{"id":11,"created_on":1516851591,"modified_on":0,"name":"312321","created_by":"4555","modified_by":"","state":1},"content":"5555","created_by":"2412","modified_by":"","state":1},"msg":"ok"}"
// @Router /api/v1/code/getCourse [post]
func GetCourse(c *gin.Context) {
	appG := app.Gin{C: c}
	courseId := com.StrTo(c.Param("courseId")).MustInt64()
	if courseId <= 0 {
		logrus.WithField("courseId",courseId).Info("invalid param")
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := &course_service.Course{CourseId: courseId}
	course, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_SERVER_ERROR, nil)
		return
	}
	if course.CourseId == 0 {
		appG.Response(http.StatusOK, e.ERROR_DATA_ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, vo.CourseVo{}.Transform(course))
}
