package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleOptions Options请求处理中间件
func HandleOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
