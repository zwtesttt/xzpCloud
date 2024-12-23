package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/user/adapters"
	"github.com/zwtesttt/xzpCloud/internal/user/app"
	"github.com/zwtesttt/xzpCloud/pkg/api/middleware"
	"github.com/zwtesttt/xzpCloud/pkg/db"
)

type Handler struct {
	*gin.Engine
	LoginHandler    *app.LoginHandler
	RegisterHandler *app.RegisterHandler
}

func New() *Handler {
	userRepo := adapters.NewUserRepository(db.GetDB())
	h := &Handler{
		Engine:          gin.New(),
		LoginHandler:    app.NewLoginHandler(userRepo),
		RegisterHandler: app.NewRegisterHandler(userRepo),
	}

	h.Use(middleware.Recovery())

	user := h.Group("/user")
	user.POST("/login", h.Login)
	user.POST("/register", h.Register)
	return h
}
