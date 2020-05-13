package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func AddArticle(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	if content == "" || title == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":http.StatusBadRequest,
			"msg":"参数错误",
			"data":"",
		})
		return
	}
	client := db.DefaultMongoClient()
	collection := client.GetCollection("blog","zsh")
	res, err :=collection.InsertOne(client.GetContext(),bson.M{"content":content})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":http.StatusInternalServerError,
			"msg":"服务器错误",
			"data":"",
		})
		errors.DebugPrintError(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":http.StatusOK,
		"msg":"",
		"data": res.InsertedID,
	})
}

