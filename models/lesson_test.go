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

func TestGetLesson(t *testing.T) {
	a, err := GetLesson(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", a)
	assert.Equal(t, int64(1), a.LessonId)
}

func TestExistLessonByID(t *testing.T) {
	c, err := ExistLessonByID(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, true, c)
}

func TestGetLessonTotal(t *testing.T) {
	c, err := GetLessonTotal(map[string]interface{}{"status": 1})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, 10, c)
}

func TestGetLessons(t *testing.T) {
	Lessons, err := GetLessons(0, 11, map[string]interface{}{"status": 1})
	if err != nil {
		t.Error(err)
		return
	}
	for _, Lesson := range Lessons {
		t.Logf("%+v", Lesson)
	}
	assert.Equal(t, 10, len(Lessons))
}

func TestAddLesson(t *testing.T) {
	Lesson := &Lesson{LessonName: "testLesson"}
	err := AddLesson(Lesson)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEditLesson(t *testing.T) {
	Lesson := &Lesson{LessonId: 20, LessonName: "testLessonsss"}
	err := AddOrUpdateClasslesson(Lesson)
	if err != nil {
		t.Error(err)
		return
	}
}
