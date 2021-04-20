package model

import "github.com/go-redis/redis"

// 初始化连接
func InitRedis() (redisclient *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "971017",
		DB:       0, // use default DB
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
