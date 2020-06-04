package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
	"github.com/masterZSH/goBlog/middlewares"
	v1 "github.com/masterZSH/goBlog/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {

	// v1版本接口
	v1Group := router.Group("/v1")
	if configs.IsDebugging() {
		v1Group.Use(middlewares.HandleOptions())
		v1Group.Use(middlewares.Cros())
	}

	{
		v1Group.GET("/user", v1.GetUser)
		v1Group.POST("/articles", v1.AddArticle)

		v1Group.GET("/articles", v1.GetArticles)
	}

}
