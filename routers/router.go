package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
	"github.com/masterZSH/goBlog/middlewares"
	v1 "github.com/masterZSH/goBlog/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {

	if configs.IsDebugging() {
		router.Use(middlewares.Cros())
	}

	// options中间件  options请求处理
	router.Use(middlewares.HandleOptions())

	// Recovery中间件 任何panic会写入500
	router.Use(gin.Recovery())

	// Logger中间件 日志中间件
	router.Use(gin.Logger())

	// jwt
	jwt := middlewares.NewJwt()
	router.GET("/login", jwt.Login)

	// v1版本接口
	v1Group := router.Group("/v1")
	{
		v1Group.POST("/articles", jwt.IsAuth, v1.AddArticle)
		v1Group.GET("/articles", v1.GetArticles)
		v1Group.GET("/articles/:id", v1.GetArticle)
		v1Group.GET("/tags", v1.GetTags)
	}

}
