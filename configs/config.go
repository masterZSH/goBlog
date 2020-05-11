package configs

import "gopkg.in/ini.v1"

type App struct {
}

var AppConfig = &App{}

func InitConfig() {
	cfg, err = ini.Load("configs/app.ini")
	mapTo("app", AppConfig)
}
