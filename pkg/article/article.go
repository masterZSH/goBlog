package article

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Article 文章结构体
type Article struct {
	Title   string
	Author  string
	Content string
}

const (
	// ADB database
	ADB = "blog"
	// ACOLLECTION collection
	ACOLLECTION = "zsh"
	// DefaultPage 默认页数
	DefaultPage = 1
	// DefaultPageSize 默认每页数量
	DefaultPageSize = 20
)

// NewArticle 创建Article
func NewArticle(title, author, content string) *Article {
	return &Article{
		title,
		author,
		content,
	}
}

// NewArticleBson 返回article的bson类型
//  {
//    "content": "test",
//    "time": 1589447406,
//    "title": "测试标题",
//    "author": "eee"
//	}
//
func (ar *Article) NewArticleBson() bson.M {
	return bson.M{
		"title":   ar.Title,
		"author":  ar.Author,
		"content": ar.Content,
		"time":    time.Now().Unix(),
	}
}

// NewAuthorFilter 新建作者过滤
func NewAuthorFilter(author string) bson.M {
	if author == "" {
		return bson.M{}
	}
	return bson.M{
		"author": author,
	}
}

// NewSortFilter 按时间正序排序
func NewSortFilter() bson.D {
	return bson.D{
		{"time", 1},
	}
}

// NewListOpts 生成opts
func NewListOpts(page, size int) *options.FindOptions {
	skip := (page - 1) * size
	opts := options.Find().
		SetSkip(int64(skip)).
		SetSort(NewSortFilter()).
		SetLimit(int64(size))
	return opts
}
