package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
	i "github.com/masterZSH/goBlog/init"
)

func main() {
	r := gin.Default()
	i.Init(r)
	port := fmt.Sprintf(":%d", configs.AppConf.Port)
	r.Run(port)
	// r.RunTLS()
}
