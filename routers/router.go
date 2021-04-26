package routers

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.LoadHTMLGlob("./pages/**/*")

	/*
	 ****************************************************
	 *	页面
	 * **************************************************
	 */
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
	//tag
	router.GET("/home/tag", v1.Tag)
	//tag列表
	router.GET("/tag/blogList/", v1.TagList)
	//搜索列表
	router.GET("/search/blogList/", v1.KeyList)
	//登录
	router.GET("/loginbyadmin", v1.LoginPage)
	//管理首页(模板页面)
	router.GET("/admin", v1.ShowAdmin)
	//管理页面
	admin := router.Group("admin")
	{
		//管理文章页面
		admin.GET("/article", v1.AdminArticle)
		//管理分类页面
		admin.GET("/classify", v1.AdminClassify)
		//管理用户页面
		admin.GET("/user", v1.AdminUser)
		//编辑文章页面
		admin.GET("/edit", v1.EditArt)
		//添加文章页面
		admin.GET("/add", v1.AddArt)
		//编辑用户页面
		admin.GET("/editUser", v1.EditUserPage)
	}

	/*
	 ****************************************************
	 *	后台逻辑 这块只需要验证token,然后把Redis更新
	 * **************************************************
	 */
	auth := router.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//管理页面
		auth.GET("/admin", v1.ShowAdmin)
		auth.GET("/articlePage", v1.AdminArticle)
		//管理分类页面
		auth.GET("/classify", v1.AdminClassify)
		//管理用户页面
		auth.GET("/user", v1.AdminUser)
		//编辑文章页面
		auth.GET("/edit", v1.EditArt)
		//添加文章页面
		auth.GET("/add", v1.AddArt)
		//获取UserInfo
		auth.GET("user/infos", v1.GetUserInfos)

		//用户
		auth.PUT("user", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.POST("user/add", v1.AddUser)
		//修改密码
		auth.PUT("editPw", v1.EditPW)
		//分类
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
	}

	/*
	 ****************************************************
	 *	后台逻辑 先查Redis,没有的话查库,然后更新Redis
	 * **************************************************
	 */
	routers := router.Group("api/v1")
	{
		routers.GET("users", v1.GetUsers)
		routers.GET("category", v1.GetCate)
		routers.GET("articles", v1.GetArticles)
		routers.GET("cateArts/:cid", v1.GetCateArt)
		routers.GET("article/info/:id", v1.GetArtInfo)
		routers.POST("login", v1.Login)
		routers.GET("files", v1.GetFile)
		//包含tag的文章
		routers.GET("tagArts/:tag", v1.GetTagArt)
		//所有tag
		routers.GET("tags", v1.GetTags)
		//搜索关键词相关所有文章
		routers.GET("keyArts/:key", v1.GetKeyArt)
		//获取用户信息
		routers.GET("user/info", v1.GetUserInfo)
	}
	return router
}
