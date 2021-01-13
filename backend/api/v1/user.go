package v1

import (
	"blog/model"
	"blog/model/request"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

//注册
func Register(c *gin.Context)  {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := &model.User{Username:R.Username, Nickname:R.NickName, Password:R.Password, Avatar:R.HeaderImg}
	err, userReturn := service.Register(*user)
	if err == nil {
		fmt.Println("注册成功",userReturn)
	}
}
