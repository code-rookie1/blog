package service

import (
	"blog/global"
	"blog/model"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户注册
func Register(u model.User) (err error, userInter model.User){
	var user model.User
	if !errors.Is(global.GVB_DB.Where("username = ? ", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经注册"), userInter
	}
	//否则 附加uuid 密码MD5简单加密  注册
	u.UUID = uuid.NewV4()
	err = global.GVB_DB.Create(&u).Error
	return err, u
}
