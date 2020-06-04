package errors

import (
	"fmt"

	"github.com/masterZSH/goBlog/configs"
)

// DebugPrintError 调试错误
func DebugPrintError(err error) {
	if err != nil {
		// todo记录日志
		if configs.IsDebugging() {
			fmt.Printf("[GOBLOG-debug] [ERROR] %v\n", err)
		}
	}
}
