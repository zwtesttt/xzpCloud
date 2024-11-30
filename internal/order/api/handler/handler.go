package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	*gin.Engine
}

func New() *Handler {
	h := &Handler{
		Engine: gin.New(),
	}

	h.GET("/", h.Hello)
	return h
}
