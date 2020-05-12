package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
	i "github.com/masterZSH/goBlog/init"
)

func main() {
	r := gin.Default()
	i.Init(r)
	r.Run(configs.AppConf.GetPort())
	// r.RunTLS()
}
