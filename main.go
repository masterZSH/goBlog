package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/routers"
)

func main() {
	r := gin.Default()
	// 初始化路由
	routers.InitRouter(r)
	r.Run(":8080")
}