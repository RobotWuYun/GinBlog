package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
	context.HTML(http.StatusOK, "home.html", nil)
}
