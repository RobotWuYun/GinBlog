package main

import (
	"GinBlog/model"
	"GinBlog/routers"
	"GinBlog/utils"
)

func main() {
	model.InitDb()
	router := routers.InitRouter()
	//静态资源
	router.Static("../../web/", "./web")
	router.Run(utils.HttpPort)
}
