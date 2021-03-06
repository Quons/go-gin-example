package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/pkg/util"
	"github.com/Quons/go-gin-example/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

/*GetAuth 验证信息*/
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(nil, e.ERROR_AUTH_CHECK_TOKEN_FAIL)
		return
	}

	if !isExist {
		appG.Response(nil, e.ERROR_AUTH)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(nil, e.ERROR_AUTH_TOKEN)
		return
	}

	appG.Response(map[string]string{
		"token": token,
	}, e.SUCCESS)
}
