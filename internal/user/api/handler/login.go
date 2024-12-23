package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/user/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
}

func (h *Handler) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	handle, err := h.LoginHandler.Handle(ctx, &app.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	api.RenderSuccess(ctx, LoginResponse{
		Token:  handle.Token,
		UserId: handle.UserId,
		Name:   handle.Name,
		Avatar: handle.Avatar,
		Email:  req.Email,
	})

}
