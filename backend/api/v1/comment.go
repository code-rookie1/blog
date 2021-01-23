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

//添加评论
func AddComment(c *gin.Context) {
	var commentStruct model.Comment
	_ = c.ShouldBindJSON(&commentStruct)
	msg, code := utils.Validate(&commentStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	//label := &model.Label{Name: labelStruct.Name, Alias: labelStruct.Alias, Description: labelStruct.Description}
	err := service.AddComment(&commentStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败,%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

//修改评论信息
func UpdateComment(c *gin.Context) {
	var commentStruct model.Comment
	_ = c.ShouldBindJSON(&commentStruct)
	err, labelBack := service.UpdateComment(commentStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败,%v", err), c)
	} else {
		response.OkWithDetailed(labelBack, "更新成功", c)
	}
}

//获取获取列表
func GetCommentList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	msg, code := utils.Validate(&pageInfo)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, list, total := service.GetCommentList(pageInfo)
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

// 通过id获取评论
func GetCommentById(c *gin.Context) {
	var getById request.GetByID
	_ = c.ShouldBindJSON(&getById)
	msg, code := utils.Validate(&getById)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, comment := service.GetCommentById(getById.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.Comment{
			Comment: *comment,
		}, c)
	}

}

// 删除评论
func DeleteComment(c *gin.Context) {

}
