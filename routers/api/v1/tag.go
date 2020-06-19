package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/pkg/tag"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTags 获取标签列表
func GetTags(c *gin.Context) {
	client := db.DefaultMongoClient()
	collection := client.GetCollection(tag.TDB, tag.TCOLLECTION)
	tags := tag.Tags{}
	err := collection.FindOne(client.GetContext(), bson.D{}).Decode(&tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": tags.Tags,
	})
}
