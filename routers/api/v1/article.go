package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/internal/errors"
	"github.com/masterZSH/goBlog/pkg/article"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	author := c.PostForm("author")

	if content == "" || title == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
			"data": "",
		})
		return
	}

	ar := article.NewArticle(title, author, content)

	client := db.DefaultMongoClient()
	collection := client.GetCollection(article.ADB, article.ACOLLECTION)
	res, err := collection.InsertOne(client.GetContext(),
		ar.NewArticleBson())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
			"data": "",
		})
		errors.DebugPrintError(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": res.InsertedID,
	})
}
