package redis

import (
	"github.com/go-redis/redis"

	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
)

// DB redis实例
var DB *redis.Client

// Client Redis client
type Client = redis.Client

// Options ...
type Options = redis.Options

func init() {
	DB = Conn(0)
}

// ConnectDB 连接redis
func ConnectDB(opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Connect 用配置文件连接数据库
func Connect(db int, prePath string) *redis.Client {
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	password := config.GetString(prePath + "password")

	client, err := ConnectDB(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	if err != nil {
		log.GetLog().WithFields(log.Fields{"func": "redis.Connect"}).Fatal(err)
	}
	return client
}

// Conn 用配置文件的默认参数连接数据库
func Conn(db int) *redis.Client {
	return Connect(db, "redis.")
}
