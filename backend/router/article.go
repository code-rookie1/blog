package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("article")
	{
		UserRouter.POST("addArticle", v1.AddArticle)
		UserRouter.POST("updateArticle", v1.UpdateArticle)
		UserRouter.POST("getArticleList", v1.GetArticleList)
		UserRouter.POST("getArticleById", v1.GetArticleById)
	}
}
