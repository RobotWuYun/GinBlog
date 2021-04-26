package model

import (
	"GinBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" lable:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=4,max=20" lable:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" lable:"角色码"`
	Resume   string `gorm:"type:longtext" json:"resume"`
}

// todo 查询用户信息
func GetUserInfo(id string) (User, int) {
	var user User
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_ART_NOT_EXIST
	}
	return user, errmsg.SUCCSE
}

//新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username=?", name).First(&users)
	if 0 != users.ID {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

//获取用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//根据用户名查询用户信息
func GetUserByName(username interface{}) (User, int) {
	var users User
	err := db.Where("username = ?", username.(string)).First(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return users, errmsg.ERROR
	}
	return users, errmsg.SUCCSE
}

//编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["resume"] = data.Resume
	if data.Password != "" {
		maps["password"] = data.Password
	}
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ChackLogin(username string, password string) int {
	var user User
	db.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_ART_NOT_EXIST
	}
	if password != user.Password {
		return errmsg.ERROR_PASSWORD_WROING
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
