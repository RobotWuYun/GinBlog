package model

import (
	"GinBlog/utils"
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var rediserr error

// 初始化连接
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     utils.RedisHost + ":" + utils.RedisPort,
		Password: utils.RedisPassWord,
		DB:       utils.RedisDB, // use default DB
	})
	_, rediserr = RedisClient.Ping().Result()
	if rediserr != nil {
		fmt.Println("Establishing a connection:", rediserr)
	}
}
