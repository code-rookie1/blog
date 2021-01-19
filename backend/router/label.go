package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitLabelRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("label")
	{
		UserRouter.POST("addLabel", v1.AddLabel)
		UserRouter.POST("updateLabel", v1.UpdateLabel)
		UserRouter.POST("getLabelList", v1.GetLabelList)
		UserRouter.POST("getLabelById", v1.GetLabelById)
	}
}
