package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/routers/api/v1"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/getUser", v1.GetUser)

	}
}
