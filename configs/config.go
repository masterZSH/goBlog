package configs

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// AppConfig 项目配置结构
type AppConfig struct {
	Env  string
	Port int
}

// AppConf 当前项目配置
var AppConf = &AppConfig{}

// InitConfig 初始化配置
func InitConfig() {
	cfg, err := ini.Load("configs/app.ini")
	appSection := cfg.Section("app")
	appSection.MapTo(AppConf)
	gin.SetMode(AppConf.Env)
}
