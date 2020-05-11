package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/routers"
)

func main() {
	r := gin.Default()
	// 初始化路由
	routers.InitRouter(r)
    r.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
            "message": "pong",
            })
	})
	r.Run(":8080")
    r.Run() // listen and serve on 0.0.0.0:8080
}