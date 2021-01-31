package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "首页",
	})
}

func ShowBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_all_articles.html", gin.H{
		"title": "blogs",
	})
}

func ShowBlogContent(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_content.html", gin.H{
		"title": "blog",
	})
}

func ShowAboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "aboutMe.html", gin.H{
		"title": "aboutMe",
	})
}

func ShowFile(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_main.html", gin.H{
		"title": "归档",
	})
}

func ShowAdmin(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_all_no_ajax.html", gin.H{
		"title": "admin",
	})
}
