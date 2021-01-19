package request

//注册用户
type RegisterStruct struct {
	Username  string `json:"userName"  validate:"required,min=3,max=12" label:"用户名"`
	Password  string `json:"passWord"  validate:"required,min=6,max=12" label:"密码"`
	NickName  string `json:"nickName"  gorm:"default:'小蒋狗'"`
	HeaderImg string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
}

//登录用户
type LoginStruct struct {
	Username string `json:"username" validate:"required,min=3,max=12" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=12" label:"密码"`
}

//修改密码
type ChangePasswordStruct struct {
	Username    string `json:"username" validate:"required,min=3,max=12" label:"用户名"`
	Password    string `json:"password" validate:"required,min=6,max=12" label:"密码"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=12" label:"密码"`
}
