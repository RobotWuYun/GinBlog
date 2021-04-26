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

	RedisHost     string
	RedisPort     string
	RedisPassWord string
	RedisDB       int

	LoginTime int64
)

func init() { //初始化方法
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("load config err:", err)
	}

	LoadServer(file)
	LoadDataBase(file)
	LoadUser(file)
	LoadRedis(file)
	LoadTimes(file)
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

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redisParam").Key("RedisHost").MustString("127.0.0.1")
	RedisPort = file.Section("redisParam").Key("RedisPort").MustString("6379")
	RedisPassWord = file.Section("redisParam").Key("RedisPassWord").MustString("111111")
	RedisDB = file.Section("redisParam").Key("RedisDB").MustInt(0)
}

func LoadTimes(file *ini.File) {
	LoginTime = file.Section("times").Key("LoginTime").MustInt64(36000000000000)
}
