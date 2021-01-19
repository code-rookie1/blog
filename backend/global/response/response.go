package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 返回成功
func Ok(c *gin.Context) {
	Result(SUCCESS, gin.H{}, "操作成功", c)
}

// 返回成功带信息
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, gin.H{}, message, c)
}

// 返回成功带数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

// 返回成功带信息和数据
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// 返回失败
func Fail(c *gin.Context) {
	Result(ERROR, gin.H{}, "操作失败", c)
}

//返回失败带数据
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, gin.H{}, message, c)
}

// 返回失败带数据
func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
