package configs

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// AppConfig 项目配置结构
type AppConfig struct {
	Env  string
	Port int
}

// GetPort 获取端口祖父串 ":8080"
func (app *AppConfig) GetPort() string {
	return fmt.Sprintf(":%d", app.Port)
}

// AppConf 当前项目配置
var AppConf = &AppConfig{}

// InitConfig 初始化配置
func InitConfig() {
	cfg, err := ini.Load("configs/app.ini")
	loadFileError(err)
	appSection := cfg.Section("app")
	appSection.MapTo(AppConf)
	gin.SetMode(AppConf.Env)
}

func loadFileError(e error) {
	if e != nil {
		fmt.Printf("加载配置文件出错：%v\n", e)
	}
}
