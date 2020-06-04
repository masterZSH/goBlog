package configs

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// AppConfig 项目配置结构
type AppConfig struct {
	Env  string
	Port int
}

// MongoConfig mongo配置
type MongoConfig struct {
	Host    string
	Port    int
	User    string
	Pwd     string
	TimeOut int
}

// GetPort 获取端口字符串 ":8080"
func (app *AppConfig) GetPort() string {
	return fmt.Sprintf(":%d", app.Port)
}

// AppConf 当前项目配置
var AppConf = &AppConfig{}

// MongoConf mongo配置
var MongoConf = &MongoConfig{}

// InitConfig 初始化配置
func InitConfig() {
	path := os.Getenv("ConfigFilePath")
	if path == "" {
		path = "configs/app.ini"
	}
	cfg, err := ini.Load(path)
	loadFileError(err)
	appSection := cfg.Section("app")
	appSection.MapTo(AppConf)
	gin.SetMode(AppConf.Env)

	mongoSection := cfg.Section("mongo")
	mongoSection.MapTo(MongoConf)
}

func loadFileError(e error) {
	if e != nil {
		fmt.Printf("加载配置文件出错：%v\n", e)
	}
}

// IsDebugging 是否调试
func IsDebugging() bool {
	return AppConf.Env == gin.DebugMode
}

// IsReleasing 是否正式
func IsReleasing() bool {
	return AppConf.Env == gin.ReleaseMode
}

// IsTesting 是否测试
func IsTesting() bool {
	return AppConf.Env == gin.TestMode
}
