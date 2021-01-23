package initialize

import (
	"blog/middleware"
	"blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	Router.Use(middleware.Cors())
	//不验证token
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) //注册无需token的路由
	}
	//下面是注册Auth路由的
	AuthGroup := Router.Group("")
	AuthGroup.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(AuthGroup)     //注册用户路由
		router.InitLabelRouter(AuthGroup)    //注册标签路由
		router.InitCommentRouter(AuthGroup)  //注册评论路由
		router.InitCategoryRouter(AuthGroup) //注册类别路由
		router.InitArticleRouter(AuthGroup)  //注册文章路由
	}
	fmt.Println("路由注册成功！")
	return Router
}
