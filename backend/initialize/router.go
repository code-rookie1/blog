package initialize

import (
	"blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	//Router.Use()
	ApiGroup := Router.Group("")
	//下面是注册路由的
	router.InitUserRouter(ApiGroup)   //注册用户路由

	fmt.Println("路由注册成功！")
	return Router
}
