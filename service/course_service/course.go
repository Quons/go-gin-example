package course_service

import (
	"github.com/Quons/go-gin-example/models"
	log "github.com/sirupsen/logrus"
)

type Course struct {
	CourseID int64 `json:"courseId"`
}

func (c *Course) Get() (*models.Course, error) {
	course, err := models.GetCourse(c.CourseID)
	if err != nil {
		log.WithField("courseId", c.CourseID).Error(err)
		return course, err
	}
	return course, nil
}
