package routers

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.LoadHTMLGlob("./pages/**/*")

	//页面
	//首页
	router.GET("/home/index", v1.ShowIndex)
	//文章列表
	router.GET("/home/blog", v1.ShowBlog)
	//文章详情
	router.GET("/home/blogContext", v1.ShowBlogContent)
	//关于我
	router.GET("/home/aboutMe", v1.ShowAboutMe)
	//归档
	router.GET("/home/file", v1.ShowFile)
	//分类
	router.GET("/home/classify", v1.Classify)
	//分类列表
	router.GET("/categorys/blogList/", v1.ClassifyList)
	//登录
	router.GET("/loginbyadmin", v1.Login)
	//模板页面
	router.GET("/admin", v1.ShowAdmin)
	//管理文章页面
	router.GET("/admin/article", v1.AdminArticle)
	//管理分类页面
	router.GET("/admin/classify", v1.AdminClassify)
	//管理用户页面
	router.GET("/admin/user", v1.AdminUser)
	//编辑文章
	router.GET("/admin/edit", v1.EditArt)
	//添加文章
	router.GET("/admin/add", v1.AddArt)

	auth := router.Group("api/v1")
	{
		//管理页面

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
	return router
}
