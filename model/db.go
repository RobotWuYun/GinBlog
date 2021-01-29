package model

import (
	"GinBlog/utils"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Println("Establishing a connection:", err)
	}
	//防止自动名称复数化 要放在自动迁移前面
	db.SingularTable(true)
	//最大闲置
	db.DB().SetMaxIdleConns(10)
	//最大连接
	db.DB().SetMaxOpenConns(100)
	//最大可复用使时间 不能超过Gin框架的失效时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//自动迁移
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	//defer db.Close()

}
