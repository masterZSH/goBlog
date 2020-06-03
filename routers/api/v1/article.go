package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/internal/errors"
	"github.com/masterZSH/goBlog/pkg/article"
	"go.mongodb.org/mongo-driver/bson"
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
		"msg":  "success",
		"data": res.InsertedID,
	})
}

// GetArticles 批量获取文章
func GetArticles(c *gin.Context) {
	var page, size int
	var results []bson.M
	var err error

	page = article.DefaultPage
	size = article.DefaultPageSize

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
	}

	if c.Query("size") != "" {
		size, err = strconv.Atoi(c.Query("size"))
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	author := c.Query("author")
	filter := article.NewAuthorFilter(author)
	client := db.DefaultMongoClient()
	collection := client.GetCollection(article.ADB, article.ACOLLECTION)
	opts := article.NewListOpts(page, size)
	res, err := collection.Find(client.GetContext(), filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "数据库连接错误",
			"data": "",
		})
		return
	}

	if err = res.All(client.GetContext(), &results); err != nil {
		errors.DebugPrintError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
			"data": "",
		})
		return
	}

	// 跨域处理
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Request-Method", "GET,POST,PUT,POST")

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers:x-requested-with", "x-requested-with,content-type")

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": gin.H{
			"page": page,
			"size": size,
			"list": results,
		},
	})
}
