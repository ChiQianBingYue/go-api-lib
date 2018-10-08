package mongo

import (
	"context"
	"fmt"

	"github.com/ChiQianBingYue/go-api-lib/config"
	"github.com/ChiQianBingYue/go-api-lib/log"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Client 默认数据库实例
var db *mongo.Database

// Init 默认初始化
func Init() {
	db = Conn()
}

// ConnectDB 用连接字符串连接数据库
func ConnectDB(dataSourceName string) (*mongo.Client, error) {
	client, err := mongo.NewClient(dataSourceName)
	// client, err := mongo.NewClient("mongodb://foo:bar@localhost:27017")
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Connect 用配置文件连接数据库
func Connect(prePath string) *mongo.Database {
	user := config.GetString(prePath + "user")
	password := config.GetString(prePath + "password")
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	dbname := config.GetString(prePath + "dbname")
	// connectTimeout := config.GetString(prePath + "connectTimeout")

	dataSourceName := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	client, err := ConnectDB(dataSourceName)
	db := client.Database(dbname)
	// db, err := ConnectDB("mongodb://" + user + "foo:bar@localhost:27017")
	if err != nil {
		log.GetLog().WithFields(log.Fields{"func": "mongo.Connect"}).Fatal(err)
	}
	return db
}

// Conn 用配置文件的默认参数连接数据库
func Conn() *mongo.Database {
	return Connect("mongo.options.")
}
