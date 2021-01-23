package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("category")
	{
		UserRouter.POST("addCategory", v1.AddCategory)
		UserRouter.POST("updateCategory", v1.UpdateCategory)
		UserRouter.POST("getCategoryList", v1.GetCategoryList)
		UserRouter.POST("getCategoryById", v1.GetCategoryById)
	}
}
