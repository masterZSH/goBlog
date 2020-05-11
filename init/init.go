package init

import (
	"github.com/masterZSH/goBlog/routers"
)

// Init 项目整体初始化
func Init() {
	// 初始化配置
	
	// 初始化路由
	routers.InitRouter(r)
	
}
