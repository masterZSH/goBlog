package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masterZSH/goBlog/internal/db"
	"github.com/masterZSH/goBlog/internal/errors"
	"github.com/masterZSH/goBlog/pkg/article"
	"github.com/masterZSH/goBlog/pkg/tag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 业务较简单 就放这里了

// Article  绑定前端传入的Article结构
type Article struct {
	Content string   `from:"content"`
	Title   string   `from:"title"`
	Author  string   `from:"author"`
	Tags    []string `from:"tags"`
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var art Article
	if err := c.BindJSON(&art); err != nil {
		return
	}
	if art.Content == "" || art.Title == "" || art.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
			"data": "",
		})
		return
	}
	ar := article.NewArticle(art.Title, art.Author, art.Content, art.Tags)
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

	// 添加标签
	{
		collection = client.GetCollection(tag.TDB, tag.TCOLLECTION)
		tags := tag.Tags{}
		tagsMap := make(map[string]bool)
		err = collection.FindOne(client.GetContext(), bson.D{}).Decode(&tags)
		// 文档为空
		if err == mongo.ErrNoDocuments {
			tagsBson := tag.NewTagsBson(ar.Tags)
			res, err = collection.InsertOne(client.GetContext(), tagsBson)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
					"data": nil,
				})
				return
			}
		}

		for _, item := range tags.Tags {
			tagsMap[item] = true
		}
		for _, item := range ar.Tags {
			if _, ok := tagsMap[item]; !ok {
				tagsMap[item] = true
			}
		}
		var tagKeys []string
		for k := range tagsMap {
			tagKeys = append(tagKeys, k)
		}
		update := bson.D{{"$set", bson.D{{"tags", tagKeys}}}}

		updateRes := collection.FindOneAndUpdate(client.GetContext(), bson.D{}, update)
		if updateRes.Err() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  updateRes.Err().Error(),
				"data": nil,
			})
			return
		}
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
	if c.Query("tag") != "" {
		filter["tags"] = bson.M{
			"$in": []string{c.Query("tag")},
		}
	}

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

// GetArticle 获取
func GetArticle(c *gin.Context) {
	var result bson.M
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误",
			"data": "",
		})
		return
	}
	// filter := article.NewArticleFilter(id)
	client := db.DefaultMongoClient()
	collection := client.GetCollection(article.ADB, article.ACOLLECTION)
	oid, _ := primitive.ObjectIDFromHex(id)
	opts := options.FindOne().SetSort(bson.D{{"time", 1}})
	err := collection.FindOne(context.Background(), bson.D{{"_id", oid}}, opts).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": result,
	})
}
