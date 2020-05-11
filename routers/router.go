package routers

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
	}
}
