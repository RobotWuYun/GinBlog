package middleware

import (
	"GinBlog/model"
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"time"
)

//判断用户是否登录
func IsUserLogin(username string) (exist bool, code int) {
	exists, err := model.RedisClient.Do("EXISTS", username).Bool()
	if err != nil {
		return false, errmsg.ERROR
	}
	return exists, errmsg.SUCCSE
}

//设置用户名缓存
func SetUserLogin(username string) int {
	err := model.RedisClient.Set(username, username, time.Duration(utils.LoginTime)).Err()
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
