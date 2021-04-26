package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int

	c.ShouldBindJSON(&data)
	//调用redis判断用户是否登录
	exists, code := middleware.IsUserLogin(data.Username)
	if code == errmsg.SUCCSE {
		if exists {
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR_USER_LOGINING,
				"message": errmsg.GetErrMsg(errmsg.ERROR_USER_LOGINING),
				"token":   "",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_REIDS,
			"message": errmsg.GetErrMsg(code),
			"token":   "",
		})
		return
	}
	//验证密码
	code = model.ChackLogin(data.Username, data.Password)

	if code == errmsg.SUCCSE {
		code = middleware.SetUserLogin(data.Username)
		if code == errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
				"token":   "",
			})
			return
		}
		token, _ = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
