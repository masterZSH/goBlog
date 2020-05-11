package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/routers"
)

func main() {
	r := gin.Default()
	
	r.Run(":8080")
}