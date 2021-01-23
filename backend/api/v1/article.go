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

//添加文章
func AddArticle(c *gin.Context) {
	var articleStruct model.Article
	_ = c.ShouldBindJSON(&articleStruct)
	msg, code := utils.Validate(&articleStruct)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	//label := &model.Label{Name: labelStruct.Name, Alias: labelStruct.Alias, Description: labelStruct.Description}
	err := service.AddArticle(&articleStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败,%v", err), c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

//修改类别信息
func UpdateArticle(c *gin.Context) {
	var articleStruct model.Article
	_ = c.ShouldBindJSON(&articleStruct)
	err, CategoryBack := service.UpdateArticle(articleStruct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败,%v", err), c)
	} else {
		response.OkWithDetailed(CategoryBack, "更新成功", c)
	}
}

//获取获取列表
func GetArticleList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	msg, code := utils.Validate(&pageInfo)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, list, total := service.GetArticleList(pageInfo)
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
func GetArticleById(c *gin.Context) {
	var getById request.GetByID
	_ = c.ShouldBindJSON(&getById)
	msg, code := utils.Validate(&getById)
	if code == response.ERROR {
		response.FailWithMessage(msg, c)
		return
	}
	err, article := service.GetArticleById(getById.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.OkWithData(resp.Article{
			Article: *article,
		}, c)
	}

}

// 删除类别
func DeleteArticle(c *gin.Context) {

}
