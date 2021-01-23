package v1

import (
	"blog/global/response"
	"blog/model"
	"blog/model/request"
	resp "blog/model/response"
	"blog/service"
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//添加类别
func AddCategory(c *gin.Context) {
	var categoryStruct model.Category
	_ = c.ShouldBindJSON(&categoryStruct)
	msg, code := utils.Validate(&categoryStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	//label := &model.Label{Name: labelStruct.Name, Alias: labelStruct.Alias, Description: labelStruct.Description}
	err := service.AddCategory(&categoryStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败,%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

//修改类别信息
func UpdateCategory(c *gin.Context) {
	var CategoryStruct model.Category
	_ = c.ShouldBindJSON(&CategoryStruct)
	err, CategoryBack := service.UpdateCategory(CategoryStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败,%v", err), c)
	} else {
		response.OkWithDetailed(CategoryBack, "更新成功", c)
	}
}

//获取获取列表
func GetCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	msg, code := utils.Validate(&pageInfo)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, list, total := service.GetCategoryList(pageInfo)
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

// 通过id获取类别
func GetCategoryById(c *gin.Context) {
	var getById request.GetByID
	_ = c.ShouldBindJSON(&getById)
	msg, code := utils.Validate(&getById)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, category := service.GetCategoryById(getById.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.Category{
			Category: *category,
		}, c)
	}

}

// 删除类别
func DeleteCategory(c *gin.Context) {

}
