package tag

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// TDB database
	TDB = "blog"
	// TCOLLECTION collection
	TCOLLECTION = "tag"
	// PoolTimeOut 对象池超时时间
	PoolTimeOut = 60 * time.Second
)

// PoolTags 对象池标签
type PoolTags struct {
	TagItem Tags
	Time time.Time
}

// Tags 标签结构体
type Tags struct {
	ID   primitive.ObjectID `json:"-" bson:"_id"`
	Tags []string           `json:"tags"`
}

// NewTagsBson tags
func NewTagsBson(tags []string) bson.M {
	return bson.M{
		"tags": tags,
	}
}

//IsEmpty 判断是否为空
func (t *PoolTags) IsEmpty() bool {
	return reflect.DeepEqual(t, &PoolTags{})
}
