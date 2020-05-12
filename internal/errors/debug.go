package errors

import (
	"fmt"

	"github.com/masterZSH/goBlog/configs"
)

const debugMode = "debug"
const onlineMode = "online"

func isDebugging() bool {
	return configs.AppConf.Env == debugMode
}

func isOnline() bool {
	return configs.AppConf.Env == onlineMode
}

func debugPrintError(err error) {
	if err != nil {
		// todo记录日志
		if isDebugging() {
			fmt.Printf("[GOBLOG-debug] [ERROR] %v\n", err)
		}
	}
}
