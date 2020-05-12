package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/configs"
)

func isDebugging() bool {
	return configs.AppConf.Env == gin.DebugMode
}

func isReleasing() bool {
	return configs.AppConf.Env == gin.ReleaseMode
}

func isTesting() bool {
	return configs.AppConf.Env == gin.TestMode
}

func debugPrintError(err error) {
	if err != nil {
		// todo记录日志
		if isDebugging() {
			fmt.Printf("[GOBLOG-debug] [ERROR] %v\n", err)
		}
	}
}
