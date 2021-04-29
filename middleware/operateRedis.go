package middleware

import (
	"GinBlog/model"
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//存储response body的结构体
type GinContext struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w GinContext) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

//判断key是否
func IsKeyExists(key string) (exist bool, code int) {
	exists, err := model.RedisClient.Do("EXISTS", key).Bool()
	if err != nil {
		return false, errmsg.ERROR
	}
	return exists, errmsg.SUCCSE
}

//设置用户名缓存
func SetKey(key string, value string, time time.Duration) int {
	err := model.RedisClient.Set(key, value, time).Err()
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//redis请求中间件
func RedisStringData() gin.HandlerFunc {
	return func(c *gin.Context) {
		//前置处理(redis中有没有,有,直接返回;没有,继续执行)
		url := c.Request.RequestURI
		blw := &GinContext{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		exists, code := IsKeyExists(url)
		if code == errmsg.SUCCSE {
			if exists {
				jsonData := []byte(model.RedisClient.Get(url).Val())
				c.Data(http.StatusOK, "application/json", jsonData)
				c.Abort()
				return
			} else {
				blw := &GinContext{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
				c.Writer = blw
				c.Next()
				//后置处理(执行完重写redis)
				SetKey(url, blw.body.String(), time.Duration(utils.LoginTime))
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR_REIDS,
				"message": errmsg.GetErrMsg(code),
				"token":   "",
			})
			return
		}
	}
}
