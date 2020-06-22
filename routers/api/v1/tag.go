package v1

import (
	"context"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/pkg/tag"
	"go.mongodb.org/mongo-driver/bson"
)

var pool = &sync.Pool{
	New: func() interface{} {
		return new(tag.PoolTags)
	},
}

// GetTags 获取标签列表
func GetTags(c *gin.Context) {
	tags := pool.Get().(*tag.PoolTags)
	if !tags.IsEmpty() && time.Now().After(tags.Time) {
		tags = pool.Get().(*tag.PoolTags)
		runtime.GC()
	}
	if !tags.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": tags.TagItem.Tags,
		})
		pool.Put(tags)
		return
	}
	client := db.DefaultMongoClient()
	collection := client.GetCollection(tag.TDB, tag.TCOLLECTION)
	err := collection.FindOne(context.Background(), bson.D{}).Decode(&tags.TagItem)
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
		"data": tags.TagItem.Tags,
	})
	tags.Time = time.Now().Add(tag.PoolTimeOut)
	pool.Put(tags)
}
