package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser 获取用户信息
func GetUser(c *gin.Context) {
	userID := c.Query("user_id")
	if userID != "" {
		c.JSON(http.StatusOK, gin.H{
			"name": "test",
			"age":  12,
		})
	}
	if userID == "" {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}
