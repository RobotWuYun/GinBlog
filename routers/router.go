package routers

import (
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	{
		//用户
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "success",
			})
		})
	}
	r.Run(utils.HttpPort)
}
