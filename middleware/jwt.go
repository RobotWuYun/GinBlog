package middleware

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var code int

//生成tocken
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Duration(utils.LoginTime))
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	tocken, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return tocken, errmsg.SUCCSE
}

//验证tocken
func CheckToken(tocken string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(tocken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}
}

//从token获取用户信息
func GetUserInfo(c *gin.Context) (jwt.MapClaims, int) {
	tockenHeader := c.Request.Header.Get("Authorization")
	checkToken := strings.SplitN(tockenHeader, " ", 2)
	tokenInfo, _ := jwt.Parse(checkToken[1], func(token *jwt.Token) (i interface{}, e error) {
		return []byte(utils.JwtKey), nil
	})
	err := tokenInfo.Claims.Valid()
	if err != nil {
		return nil, errmsg.ERROR
	} else {
		finToken := tokenInfo.Claims.(jwt.MapClaims)
		//校验下token是否过期
		succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
		if succ != true {
			return nil, errmsg.ERROR_TOCKEN_RUNTIME
		} else {
			return finToken, errmsg.SUCCSE
		}
	}
}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tockenHeader := c.Request.Header.Get("Authorization")
		code = errmsg.SUCCSE
		if tockenHeader == "" {
			code = errmsg.ERROR_TOCKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tockenHeader, " ", 2)
		if len(checkToken) == 2 && checkToken[0] == "Bearer" && checkToken[1] == "null" {
			code = errmsg.ERROR_TOCKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOCKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOCKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOCKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()

	}
}
