package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"gorm.io/gorm"
)

// 增加评论
func AddComment(comment *model.Comment) (err error) {
	var commentStruct model.Comment
	if !errors.Is(global.GVB_DB.Where("id = ? ", comment.ID).First(&commentStruct).Error, gorm.ErrRecordNotFound) {
		return errors.New("评论id重复")
	}
	err = global.GVB_DB.Create(&comment).Error
	return err
}

// 修改评论
func UpdateComment(comment model.Comment) (err error, commentBack *model.Comment) {
	var commentStruct model.Comment
	err = global.GVB_DB.Where("id = ?", comment.ID).First(&commentStruct).Updates(&comment).Error
	return err, &commentStruct
}

// 获取评论列表
func GetCommentList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVB_DB.Model(&model.Comment{})
	var commentList []model.Comment
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&commentList).Error
	return err, commentList, total
}

// 通过id获取评论
func GetCommentById(id int) (err error, commentBack *model.Comment) {
	var commentStruct model.Comment
	err = global.GVB_DB.Where("id = ?", id).First(&commentStruct).Error
	return err, &commentStruct
}
