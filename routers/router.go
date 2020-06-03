package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/masterZSH/goBlog/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {

	// todo debug版本跨域处理中间见

	// v1版本接口
	v1Group := router.Group("/v1")
	{
		v1Group.GET("/user", v1.GetUser)
		v1Group.POST("/articles", v1.AddArticle)

		v1Group.GET("/articles", v1.GetArticles)
		v1Group.OPTIONS("/articles", v1.GetArticles)
	}

}
