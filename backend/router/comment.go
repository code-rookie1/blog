package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("comment")
	{
		UserRouter.POST("addComment", v1.AddComment)
		UserRouter.POST("updateComment", v1.UpdateComment)
		UserRouter.POST("getCommentList", v1.GetCommentList)
		UserRouter.POST("getCommentById", v1.GetCommentById)
	}
}
