package routers

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	if usedrId := c.PostForm("user_id"); usedrId != "" {
		c.JSON(200, gin.H{
			"user": {
				"name": "test",
				"age": 12
			},
		})
	}
}
