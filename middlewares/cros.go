package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cros 跨域处理中间件 
func Cros() gin.HandlerFunc{
	return func (c *gin.Context)  {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Header("Access-Control-Request-Method", "GET,POST,PUT,POST,DELETE,OPTIONS")
		c.Next()
	}
}