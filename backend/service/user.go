package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户注册
func Register(user model.User) (err error, userInter model.User) {
	var UserStruct model.User
	if !errors.Is(global.GVB_DB.Where("username = ? ", user.Username).First(&UserStruct).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经注册"), userInter
	}
	//否则 附加uuid 密码MD5简单加密  注册
	user.UUID = uuid.NewV4()
	err = global.GVB_DB.Create(&user).Error
	return err, userInter
}

// 登录
func Login(user *model.User) (err error, userInter *model.User) {
	var UserStruct model.User
	err = global.GVB_DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&UserStruct).Error
	return err, &UserStruct
}

// 修改个人信息
func SetUserInfo(user model.User) (err error, userInter model.User) {
	var UserStruct model.User
	err = global.GVB_DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&UserStruct).Updates(&user).Error
	return err, userInter
}

//修改密码
func ChangePassword(user *model.User, newPassword string) (err error, userInter *model.User) {
	var userStruct model.User
	err = global.GVB_DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&userStruct).Update("password", newPassword).Error
	return err, userInter
}

// 获取用户列表
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVB_DB.Model(&model.User{})
	var userList []model.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

// 通过id获取用户
func GetUserById(id int) (err error, userInter *model.User) {
	var userStruct model.User
	err = global.GVB_DB.Where("id = ?", id).First(&userStruct).Error
	return err, &userStruct
}
