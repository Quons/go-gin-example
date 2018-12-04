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

func TestGetClassLesson(t *testing.T) {
	a, err := GetLesson(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", a)
	assert.Equal(t, int64(1), a.LessonId)
}

func TestGetClasslessonsByPhase(t *testing.T) {
	classLessons, err := GetClasslessonsByPhase(2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", classLessons)
	for _, value := range classLessons {
		t.Logf("%+v", value.Lesson)
		t.Logf("startTime:%v",value.StartTime)
	}
}

func TestExistClassLessonByID(t *testing.T) {
	c, err := ExistLessonByID(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, true, c)
}

func TestGetClassLessonTotal(t *testing.T) {
	c, err := GetLessonTotal(map[string]interface{}{"status": 1})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", c)
	assert.Equal(t, 10, c)
}

func TestGetClassLessons(t *testing.T) {
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

func TestAddClassLesson(t *testing.T) {
	Lesson := &Lesson{LessonName: "testLesson"}
	err := AddLesson(Lesson)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEditClassLesson(t *testing.T) {
	Lesson := Lesson{LessonId: 20, LessonName: "testLessonsss"}
	err := EditLesson(Lesson)
	if err != nil {
		t.Error(err)
		return
	}
}