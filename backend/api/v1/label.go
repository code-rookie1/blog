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

//添加标签
func AddLabel(c *gin.Context) {
	var labelStruct model.Label
	_ = c.ShouldBindJSON(&labelStruct)
	msg, code := utils.Validate(&labelStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	label := &model.Label{Name: labelStruct.Name, Alias: labelStruct.Alias, Description: labelStruct.Description}
	err := service.AddLabel(label)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败,%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

//修改标签信息
func UpdateLabel(c *gin.Context) {
	var labelStruct model.Label
	_ = c.ShouldBindJSON(&labelStruct)
	err, labelBack := service.UpdateLabel(labelStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败,%v", err), c)
	} else {
		response.OkWithDetailed(labelBack, "更新成功", c)
	}
}

//获取获取列表
func GetLabelList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	msg, code := utils.Validate(&pageInfo)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, list, total := service.GetLabelList(pageInfo)
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

// 通过id获取标签
func GetLabelById(c *gin.Context) {
	var getById request.GetByID
	_ = c.ShouldBindJSON(&getById)
	msg, code := utils.Validate(&getById)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, label := service.GetLabelById(getById.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.Label{
			Label: *label,
		}, c)
	}

}

// 删除标签
func DeleteLabel(c *gin.Context) {

}
