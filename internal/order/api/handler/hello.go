package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Hello(context *gin.Context) {
	context.String(200, "Hello World")
}
