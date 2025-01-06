package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 成功响应（带数据）
func RenderSuccess(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Msg:  "success",
		Data: body,
	})
}

// 成功响应（无数据）
func RenderSuccessNoBody(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

// 成功响应（带自定义消息和状态码）
func RenderSuccessWithMsgData(c *gin.Context, body interface{}, msg string, code int) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: body,
	})
}

// 错误响应（默认状态码为 400）
func RenderError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Code: FindErrorsCode(err),
		Msg:  err.Error(),
	})
}

// 错误响应（带自定义状态码）
func RenderErrorWithStatus(c *gin.Context, err error, code int) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: code,
		Msg:  err.Error(),
	})
}

// 未认证
func RenderUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: Unauthorized,
		Msg:  StrUnauthorized,
	})
}

func RenderInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
	})
}

// 错误响应（参数错误）
func RenderBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  "invalid parameter",
	})
}
