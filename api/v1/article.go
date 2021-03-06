package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//todo 添加文章
func AddArticle(c *gin.Context) {

	var data model.Article

	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// todo 查询分类下所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"total":    total,
		"pagesize": pageSize,
		"pagenum":  pageNum,
		"message":  errmsg.GetErrMsg(code),
	})
}

// todo 查询包含tag的所有文章
func GetTagArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	tag := c.Param("tag")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetTagArt(tag, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"total":    total,
		"pagesize": pageSize,
		"pagenum":  pageNum,
		"message":  errmsg.GetErrMsg(code),
	})
}

// todo 查询单个文章信息
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"total":    total,
		"pagesize": pageSize,
		"pagenum":  pageNum,
		"message":  errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询包含key的所有文章
func GetKeyArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	key := c.Param("key")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetKeyArt(key, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"data":     data,
		"total":    total,
		"pagesize": pageSize,
		"pagenum":  pageNum,
		"message":  errmsg.GetErrMsg(code),
	})
}

// todo 查询所有tag
func GetTags(c *gin.Context) {
	data, code, _ := model.GetArt(-1, -1)
	tagMap := make(map[string]int)
	var tags []string
	for _, val := range data {
		tags = strings.Split(val.Tag, ",")
		for _, val := range tags {
			num, ok := tagMap[val]
			if ok {
				tagMap[val] = num + 1
			} else {
				tagMap[val] = 1
			}
		}
	}
	delete(tagMap, "")
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    tagMap,
		"total":   len(tagMap),
		"message": errmsg.GetErrMsg(code),
	})
}
