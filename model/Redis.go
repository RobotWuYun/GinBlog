package model

import (
	"GinBlog/utils"
	"github.com/go-redis/redis"
)

// 初始化连接
func InitRedis() (redisclient *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.RedisHost + ":" + utils.RedisPort,
		Password: utils.DbPassWord,
		DB:       utils.RedisDB, // use default DB
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
