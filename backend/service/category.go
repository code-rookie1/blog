package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"gorm.io/gorm"
)

// 增加类别
func AddCategory(category *model.Category) (err error) {
	var categoryStruct model.Category
	if !errors.Is(global.GVB_DB.Where("name = ? ", category.Name).First(&categoryStruct).Error, gorm.ErrRecordNotFound) {
		return errors.New("标签名字重复")
	}
	err = global.GVB_DB.Create(&category).Error
	return err
}

// 修改类别
func UpdateCategory(category model.Category) (err error, categoryBack *model.Category) {
	var categoryStruct model.Category
	err = global.GVB_DB.Where("id = ?", category.ID).First(&categoryStruct).Updates(&category).Error
	return err, &categoryStruct
}

// 获取类别列表
func GetCategoryList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVB_DB.Model(&model.Category{})
	var categoryList []model.Category
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&categoryList).Error
	return err, categoryList, total
}

// 通过id获取类别
func GetCategoryById(id int) (err error, categoryBack *model.Category) {
	var categoryStruct model.Category
	err = global.GVB_DB.Where("id = ?", id).Preload("Article").Preload("SubCategory").First(&categoryStruct).Error
	return err, &categoryStruct
}
