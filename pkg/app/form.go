package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/sirupsen/logrus"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		logrus.Errorf(err.Error())
		return http.StatusOK, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logrus.Errorf(err.Error())
		return http.StatusOK, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusOK, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
