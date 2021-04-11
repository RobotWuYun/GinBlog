package model

import (
	"GinBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null" json:"title"`
	Cid       int    `gorm:"type:int;not null" json:"cid"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	Content   string `gorm:"type:longtext" json:"content"`
	Tag       string `gorm:"type:varchar(100)" json:"tag"`
	Timestamp string `gorm:"type:varchar(10)" json:"timestamp"`

	Category Category `gorm:"foreignkey:Cid"`
}

//新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// todo 查询分类下所有文章
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtiList []Article
	var total int
	err := db.Order("created_at desc").Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", cid).Find(&cateArtiList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtiList, errmsg.SUCCSE, total
}

// todo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

//查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Preload("Category").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

//查询文章列表
func GetArtAsc(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Order("created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Preload("Category").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

//查询有tag的文章列表
func GetTagArt(tag string, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Order("created_at desc").Where("tag LIKE ?", "%"+tag+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Preload("Category").Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

//关键字查询-文章列表
func GetKeyArt(key string, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	err := db.Order("created_at desc").Where("title LIKE ?", "%"+key+"%").Or("tag LIKE ?", "%"+key+"%").Or("`desc` LIKE ?", "%"+key+"%").Or("`timestamp` LIKE ?", "%"+key+"%").Or("content LIKE ?", "%"+key+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Preload("Category").Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

//编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["tag"] = data.Tag
	maps["timestamp"] = data.Timestamp
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
