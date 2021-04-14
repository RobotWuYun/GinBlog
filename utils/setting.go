package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//做数据处理

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Id string
)

func init() { //初始化方法
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("load config err:", err)
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadUser(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("sfd3SGFiofs2432jsFS")
}

func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHort").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadUser(file *ini.File) {
	Id = file.Section("user").Key("Id").MustString("1")
}
