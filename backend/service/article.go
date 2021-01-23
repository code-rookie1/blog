package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"gorm.io/gorm"
)

// 增加文章
func AddArticle(article *model.Article) (err error) {
	var articleStruct model.Article
	if !errors.Is(global.GVB_DB.Where("title = ? ", article.Title).First(&articleStruct).Error, gorm.ErrRecordNotFound) {
		return errors.New("文章名字重复")
	}
	err = global.GVB_DB.Create(&article).Error
	return err
}

// 修改文章
func UpdateArticle(article model.Article) (err error, articleBack *model.Article) {
	var articleStruct model.Article
	err = global.GVB_DB.Where("id = ?", article.ID).First(&articleStruct).Updates(&article).Error
	return err, &articleStruct
}

// 获取文章列表
func GetArticleList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVB_DB.Model(&model.Article{})
	var articleList []model.Article
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	return err, articleList, total
}

// 通过id获取文章
func GetArticleById(id int) (err error, articleBack *model.Article) {
	var articleStruct model.Article
	err = global.GVB_DB.Where("id = ?", id).First(&articleStruct).Error
	return err, &articleStruct
}
