package db
import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	i "github.com/masterZSH/goBlog/init"
	"go.mongodb.org/mongo-driver/bson"
)
func TestMongo(t *testing.T){
	r := gin.Default()
	os.Setenv("ConfigFilePath","../../configs/app.ini")
	i.Init(r)
	client := DefaultMongoClient()

	collection := client.GetCollection("zsh","blog")

	// 新增一条
	res, err := collection.InsertOne(client.GetContext(), bson.M{"name": "pi", "value": 3.14159})
	
	// 第二条
	res, err = collection.InsertOne(client.GetContext(), bson.M{"name": "111", "value":123})

	id := res.InsertedID
	if id == nil{
		t.Error("新增失败")
	}
	cur, err := collection.Find(client.GetContext(), bson.D{})
	if err != nil{
		t.Error("获取失败")
	}
	defer func(){
		cur.Close(client.GetContext())
		collection.Drop(client.GetContext())
	}()
		
	for cur.Next(client.GetContext()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil{
			t.Error("记录错误")
		}
	}
}