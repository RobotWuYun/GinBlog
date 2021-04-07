package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//首页
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "首页",
	})
}

//博客列表页
func ShowBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_all_articles.html", gin.H{
		"title": "blogs",
	})
}

//博客详情页
func ShowBlogContent(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_content.html", gin.H{
		"title": "blog",
	})
}

//关于我页面
func ShowAboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "aboutMe.html", gin.H{
		"title": "aboutMe",
	})
}

//归档页面
func ShowFile(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_main.html", gin.H{
		"title": "归档",
	})
}

//管理页面
func ShowAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_all_no_ajax.html", gin.H{
		"title": "后台管理",
	})
}

//管理文章页面
func AdminArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_article.html", gin.H{
		"title": "文章管理",
	})
}

//管理分类页面
func AdminClassify(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_tag.html", gin.H{
		"title": "分类管理",
	})
}

//管理用户页面
func AdminUser(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_user.html", gin.H{
		"title": "用户管理",
	})
}

//登录页面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "userLogin.html", gin.H{
		"title": "login",
	})
}

//分类页面
func Classify(c *gin.Context) {
	c.HTML(http.StatusOK, "Classify.html", gin.H{
		"title": "分类",
	})
}

//博客分类列表页
func ClassifyList(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_category_articles.html", gin.H{
		"title": "分类文章列表",
	})
}

//编辑文章
func EditArt(c *gin.Context) {
	c.HTML(http.StatusOK, "editArts.html", gin.H{
		"title": "编辑文章",
	})
}

//添加文章
func AddArt(c *gin.Context) {
	c.HTML(http.StatusOK, "addArts.html", gin.H{
		"title": "添加文章",
	})
}
