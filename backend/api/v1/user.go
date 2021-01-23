package v1

import (
	"blog/global"
	"blog/global/response"
	"blog/middleware"
	"blog/model"
	"blog/model/request"
	resp "blog/model/response"
	"blog/service"
	"blog/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

//注册
func Register(c *gin.Context) {
	var RegisterStruct request.RegisterStruct
	_ = c.ShouldBindJSON(&RegisterStruct)
	msg, code := utils.Validate(&RegisterStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	user := &model.User{Username: RegisterStruct.Username, Nickname: RegisterStruct.NickName, Password: RegisterStruct.Password, Avatar: RegisterStruct.HeaderImg}
	err := service.Register(*user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("注册失败，%v", err), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

// 登录
func Login(c *gin.Context) {
	var LoginStruct request.LoginStruct
	_ = c.ShouldBindJSON(&LoginStruct)
	msg, code := utils.Validate(&LoginStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	user := &model.User{Username: LoginStruct.Username, Password: LoginStruct.Password}
	if err, userBack := service.Login(user); err != nil {
		response.FailWithMessage("用户名或密码错误", c)
	} else {
		tokenNext(c, *userBack)
		//response.OkWithDetailed(userBack, "登录成功", c)
		//后续增加jwt
	}
}

//登录后签发token
func tokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.GVB_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		NickName:   user.Nickname,
		Username:   user.Username,
		BufferTime: global.GVB_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVB_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "UltraMan",                                            // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	} else {
		response.OkWithDetailed(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}

}

//修改个人信息
func SetUserInfo(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	msg, code := utils.Validate(&user)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, userBack := service.SetUserInfo(user)
	if err != nil {
		fmt.Println("更新失败", err)
		response.FailWithMessage(fmt.Sprintf("更新失败,%v", err), c)
	} else {
		response.OkWithDetailed(userBack, "更新成功", c)
	}
}

//修改密码
func ChangePassword(c *gin.Context) {
	var ChangePasswordStruct request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&ChangePasswordStruct)
	msg, code := utils.Validate(&ChangePasswordStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	user := &model.User{Username: ChangePasswordStruct.Username, Password: ChangePasswordStruct.Password}
	if err, _ := service.ChangePassword(user, ChangePasswordStruct.NewPassword); err != nil {
		response.FailWithMessage("修改密码失败", c)
		return
	} else {
		response.OkWithMessage("修改密码成功", c)
	}

}

//获取用户列表
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	msg, code := utils.Validate(&pageInfo)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// 通过id获取用户
func GetUserById(c *gin.Context) {
	var getById request.GetByID
	_ = c.ShouldBindJSON(&getById)
	msg, code := utils.Validate(&getById)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, user := service.GetUserById(getById.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.User{
			User: *user,
		}, c)
	}

}

// 删除用户
func DeleteUser(c *gin.Context) {

}
