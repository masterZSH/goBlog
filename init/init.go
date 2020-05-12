package init

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
	"github.com/masterZSH/goBlog/routers"
)

// Init 项目整体初始化
func Init(r *gin.Engine) {
	// 初始化项目配置
	configs.InitConfig()
	// 初始化路由
	routers.InitRouter(r)

}
