package routers

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	{
		//用户
		//auth.POST("user/add", v1.AddUser)
		auth.GET("users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.POST("user/add", v1.AddUser)
		//分类
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		auth.GET("categorys", v1.GetCate)
		//文章
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		auth.GET("articles", v1.GetArticles)
		auth.GET("articles/list/:cid", v1.GetCateArt)
		auth.GET("article/info/:id", v1.GetArtInfo)

	}
	r.Run(utils.HttpPort)
}
