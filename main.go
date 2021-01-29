package main

import (
	"GinBlog/model"
	"GinBlog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
