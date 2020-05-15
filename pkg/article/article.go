package article

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
)

// NewArticle 创建Article
func NewArticle(title, author, content string) *Article {
	return &Article{
		title,
		author,
		content,
	}
}

// GetBson 返回article的bson类型
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

func  NewAuthorFilter(author string) bson.M{
	return bson.M{
		"title":   ar.Title,
		"author":  ar.Author,
		"content": ar.Content,
		"time":    time.Now().Unix(),
	}
}
