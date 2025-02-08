package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/vm/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

func (h *Handler) DeleteVm(ctx *gin.Context) {
	//TODO userId
	id := ctx.Param("id")
	if id == "" {
		api.RenderBadRequest(ctx)
		return
	}
	if err := h.deleteVmHandler.Handle(ctx, &app.DeleteVmInput{Id: id, UserId: ""}); err != nil {
		api.RenderError(ctx, err)
		return
	}

	api.RenderSuccessNoBody(ctx)
}
