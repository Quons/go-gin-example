package models

import (
	"github.com/Quons/go-gin-example/pkg/setting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	setting.Setup("dev")
	Setup()
}

func TestGetLessonSection(t *testing.T) {
	a, err := GetLessonSection(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("lessonName:%+v,courseName:%v", a.Lesson.LessonName, a.Course.CourseName)
	assert.Equal(t, int64(1), a.SectionId)
}

func TestExistLessonSectionByID(t *testing.T) {
	c, err := ExistLessonSectionByID(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, true, c)
}

func TestGetLessonSectionTotal(t *testing.T) {
	c, err := GetLessonSectionTotal(map[string]interface{}{"status": 1})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, 10, c)
}

func TestGetLessonSections(t *testing.T) {
	LessonSections, err := GetLessonSections(0, 11, map[string]interface{}{"status": 1})
	if err != nil {
		t.Error(err)
		return
	}
	for _, LessonSection := range LessonSections {
		t.Logf("%+v", LessonSection)
	}
	assert.Equal(t, 10, len(LessonSections))
}

func TestAddLessonSection(t *testing.T) {
	LessonSection := &Lessonsection{SectionName: "testLessonSection"}
	err := AddLessonSection(LessonSection)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEditLessonSection(t *testing.T) {
	LessonSection := &Lessonsection{SectionId: 20, SectionName: "testLessonSectionsss"}
	err := AddOrUpdateLessonSection(LessonSection)
	if err != nil {
		t.Error(err)
		return
	}
}