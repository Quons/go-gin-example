package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//启用swagger文档
	_ "github.com/Quons/go-gin-example/docs"
	"github.com/Quons/go-gin-example/pkg/export"
	"github.com/Quons/go-gin-example/pkg/logging"
	"github.com/Quons/go-gin-example/pkg/qrcode"
	"github.com/Quons/go-gin-example/pkg/setting"
	"github.com/Quons/go-gin-example/pkg/upload"
	"github.com/Quons/go-gin-example/routers/api"
	"github.com/Quons/go-gin-example/routers/api/v1"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"

	"github.com/Quons/go-gin-example/middleware"
	"github.com/Quons/go-gin-example/routers/api/v1/code"
)

/*路由注册*/
func registerRouter(r *gin.Engine) {
	r.GET("/auth", api.GetAuth)
	/*r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		logrus.Info(name)
		time.Sleep(20 * time.Second)
		c.String(http.StatusOK, "welcome Gin Server:%s\n", name)
	})*/
	r.POST("/upload", api.UploadImage)
	apiv1 := r.Group("/api/v1")
	//验证token
	//apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//新增文章和标签
		apiv1.GET("/articleAndTag", v1.AddArticleAndTag)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}
	apiv1Code := r.Group("/api/v1/code")
	apiv1Code.Use(middleware.CheckToken())
	{
		apiv1Code.POST("/getCourse", code.GetCourse)
	}

}

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	r := gin.New()
	//设置gin日志输出writer
	r.Use(gin.LoggerWithWriter(logging.GetGinLogWriter()))
	//设置gin恢复日志数据writer
	r.Use(gin.RecoveryWithWriter(logging.GetGinLogWriter()))

	//静态目录
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//swagger自动文档路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//注册业务路由
	registerRouter(r)
	return r
}
