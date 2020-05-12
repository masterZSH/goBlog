package db

import (
	"context"
	"fmt"
	"time"

	"github.com/masterZSH/goBlog/configs"
	"github.com/masterZSH/goBlog/internal/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient mongo客户端
type MongoClient struct {
	Client *mongo.Client
}

// NeedAuth 是否需要认证
func (mc *MongoClient) NeedAuth() bool {
	mongoConf := configs.MongoConf
	if mongoConf.User == "" {
		return false
	}
	if mongoConf.Pwd == "" {
		return false
	}
	return true
}

// CreateAuthCredential 创建认证
func (mc *MongoClient) CreateAuthCredential() options.Credential {
	authCredential := options.Credential{
		Username: configs.MongoConf.User,
		Password: configs.MongoConf.Pwd,
	}
	return authCredential
}

// GetContext 获取上下文
func (mc *MongoClient) GetContext() (ctx context.Context) {
	// context
	timeout := configs.MongoConf.TimeOut
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	return
}

// BuildClientOptions 构建客户端选项
func (mc *MongoClient) BuildClientOptions() (clientOptions *options.ClientOptions) {
	clientOptions = options.Client()
	if mc.NeedAuth() {
		clientOptions.SetAuth(mc.CreateAuthCredential())
	}
	clientOptions.ApplyURI(mc.GetMongoURI())
	return
}

// GetMongoURI 获取mongo uri
func (mc *MongoClient) GetMongoURI() string {
	return fmt.Sprintf("mongodb://%s:%d",
		configs.MongoConf.Host,
		configs.MongoConf.Port)
}

// DefaultMongoClient 默认客户端
func DefaultMongoClient() *MongoClient {
	client := new(MongoClient)

	// client
	var err error
	client.Client, err = mongo.Connect(client.GetContext(), client.BuildClientOptions())
	errors.DebugPrintError(err)
	return client
}

// GetCollection 获取collection
func (mc *MongoClient) GetCollection(db, collection string) (coll *mongo.Collection) {
	coll = mc.Client.Database(db).Collection(collection)
	return
}
