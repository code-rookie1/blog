package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"gorm.io/gorm"
)

// 增加标签
func AddLabel(label *model.Label) (err error) {
	var labelStruct model.Label
	if !errors.Is(global.GVB_DB.Where("name = ? ", label.Name).First(&labelStruct).Error, gorm.ErrRecordNotFound) {
		return errors.New("标签名字重复")
	}
	err = global.GVB_DB.Create(&label).Error
	return err
}

// 修改标签
func UpdateLabel(label model.Label) (err error, labelBack *model.Label) {
	var labelStruct model.Label
	err = global.GVB_DB.Where("name = ?", label.Name).First(&labelStruct).Updates(&label).Error
	return err, &labelStruct
}

// 获取标签列表
func GetLabelList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVB_DB.Model(&model.Label{})
	var labelList []model.Label
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&labelList).Error
	return err, labelList, total
}

// 通过id获取标签
func GetLabelById(id int) (err error, labelBack *model.Label) {
	var labelStruct model.Label
	err = global.GVB_DB.Where("id = ?", id).First(&labelStruct).Error
	return err, &labelStruct
}
