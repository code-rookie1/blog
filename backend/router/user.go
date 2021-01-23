package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("setUserInfo", v1.SetUserInfo)
		UserRouter.POST("changePassword", v1.ChangePassword)
		UserRouter.POST("getUserList", v1.GetUserList)
		UserRouter.POST("getUserById", v1.GetUserById)
	}
}
