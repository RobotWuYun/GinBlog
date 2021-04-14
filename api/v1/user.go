package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//查询用户是否存在
func UserExist() {

}

//添加用户
func AddUser(c *gin.Context) {
	//todo 添加用户
	var data model.User
	_ = c.ShouldBindJSON(&data)

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//前台查询用户简历
func GetUserInfo(c *gin.Context) {
	data, code := model.GetUserInfo(utils.Id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Resume,
		"message": errmsg.GetErrMsg(code),
	})
}

//后台获取用户信息
func GetUserInfos(c *gin.Context) {
	info, code := middleware.GetUserInfo(c)
	if code == 200 {
		data, code := model.GetUserByName(info["username"])
		userdata := make(map[string]string)
		userdata["username"] = data.Username
		userdata["content"] = data.Resume
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    userdata,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var Pagedata model.User
	c.ShouldBindJSON(&Pagedata)
	var data model.User
	var code int
	info, code := middleware.GetUserInfo(c)
	if code == 200 {
		data, code = model.GetUserByName(info["username"])
		if len(Pagedata.Username) > 0 {
			if Pagedata.Username == data.Username {
				model.EditUser(int(data.ID), &Pagedata)
			} else {
				code = model.CheckUser(Pagedata.Username)
				if code == errmsg.ERROR_USERNAME_USED {
					c.Abort()
				} else {
					model.EditUser(int(data.ID), &Pagedata)
				}
			}
		} else {
			code = errmsg.ERROR_USERNAME_NOT_NULL
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

/*
//修改密码
func EditPW(c *gin.Context) {
	c.Request.
	c.ShouldBindJSON(&Pagedata)
	var data model.User
	var code int
	info,code := middleware.GetUserInfo(c)
	if code == 200 {
		data, code = model.GetUserByName(info["username"])
		if Pagedata. {
			if Pagedata.Username == data.Username {
				model.EditUser(int(data.ID), &Pagedata)
			}else {
				code = model.CheckUser(Pagedata.Username)
				if code == errmsg.ERROR_USERNAME_USED {
					c.Abort()
				}else {
					model.EditUser(int(data.ID), &Pagedata)
				}
			}
		} else {
			code = errmsg.ERROR_USERNAME_NOT_NULL
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}*/

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
