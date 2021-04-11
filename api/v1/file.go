package v1

import (
	"GinBlog/model"
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetFile(c *gin.Context) {
	data, code, total := model.GetArtAsc(-1, -1)

	var fileMap utils.Multimap
	fileMap = make(utils.Multimap)
	var value string

	for _, val := range data {
		value = val.Title + "|||" + strconv.Itoa(int(val.ID))
		fileMap.Add(val.Timestamp, value)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    fileMap,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
