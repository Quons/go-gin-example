package v1

import (
	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/service/course_service"
	"github.com/Quons/go-gin-example/vo"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/Quons/go-gin-example/models"
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
		log.Info(err)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}
	//获取studentId
	studentId := c.GetInt64(e.PARAM_STUDENT_ID)
	studentInfo, err := models.GetStudent(studentId)
	if err != nil {
		log.WithField("studentId", studentId).Error(err)
		appG.Response(nil, e.ERROR_SERVER_ERROR)
		return
	}
	log.WithField("studentInfo", studentInfo).Info()
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
