package tag

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// TDB database
	TDB = "blog"
	// TCOLLECTION collection
	TCOLLECTION = "tag"
)

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
