package v1

import (
	"github.com/Quons/go-gin-example/pkg/app"
	"github.com/Quons/go-gin-example/pkg/e"
	"github.com/Quons/go-gin-example/service/brand_service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Tags 添加品牌
// @Summary 添加品牌
// @Description 添加品牌
// @Produce  json
// @accept application/x-www-form-urlencoded
// @Param name formData string true "品牌名称"
// @Success 99999 "ok"
// @Failure 10000 {string} json "{"code":10000,"data":{},"msg":"服务器错误"}"
// @Failure 20000 {string} json "{"code":20000,"data":{},"msg":"参数错误"}"
// @Router /api/v1/getCourse [post]
func AddBrand(c *gin.Context) {
	appG := app.Gin{C: c}
	brand := brand_service.Brand{}
	err := c.ShouldBind(&brand)
	if err != nil {
		log.Info(err)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}
	if err = brand.AddBrand(); err != nil {
		appG.Response(nil, e.ERROR_SERVER_ERROR)
		return
	}
	appG.Response("ok", e.SUCCESS)
}

// @Tags 获取品牌列表
// @Summary 获取品牌列表
// @Description 添加品牌
// @Produce  json
// @accept application/x-www-form-urlencoded
// @Success 99999 "ok"
// @Failure 10000 {string} json "{"code":10000,"data":{},"msg":"服务器错误"}"
// @Failure 20000 {string} json "{"code":20000,"data":{},"msg":"参数错误"}"
// @Router /api/v1/getCourse [post]
func GetBrandList(c *gin.Context) {
	appG := app.Gin{C: c}
	brand := brand_service.Brand{}
	list, err := brand.GetBrandList()
	if err != nil {
		log.Info(err)
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}
	appG.Response(list, e.SUCCESS)
}
