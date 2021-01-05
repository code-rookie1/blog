package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID   `json:"uuid" gorm:"comment:用户UUID"`
	Username  string      `json:"username" gorm:"comment:用户登录名"`
	Password  string      `json:"password" gorm:"comment:密码"`
	Nickname  string      `json:"nickname" gorm:"default:小蒋;comment:拥挤昵称"`
	Avatar    string      `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Email     string      `json:"email" gorm:"comment:邮箱"`
	Phone     string      `json:"phone" gorm:"comment:电话"`
	Region    string      `json:"region" gorm:"default:成都;comment:地区"`
	Age       uint         `json:"age" gorm:"default:18;comment:年龄"`
	Signature string      `json:"signature" gorm:"type:text;default:这个人很懒，什么页没留下;comment:签名"`
}