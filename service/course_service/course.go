package course_service

import (
	"github.com/Quons/go-gin-example/models"
	log "github.com/sirupsen/logrus"
)

type Course struct {
	CourseId int64
}

func (c *Course) Get() (*models.Course, error) {
	course, err := models.GetCourse(c.CourseId)
	if err != nil {
		log.WithField("courseId", c.CourseId).Error(err)
		return course, err
	}
	return course, nil
}
