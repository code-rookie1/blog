package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
	}
}
