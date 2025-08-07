package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/vm/adapters"
	"github.com/zwtesttt/xzpCloud/internal/vm/app"
	"github.com/zwtesttt/xzpCloud/pkg/api/middleware"
	"github.com/zwtesttt/xzpCloud/pkg/config"
	"github.com/zwtesttt/xzpCloud/pkg/db"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
)

type Handler struct {
	*gin.Engine
	createVmHandler *app.CreateVmHandler
	deleteVmHandler *app.DeleteVmHandler
	getVmsHandler   *app.GetVmsHandler
	startVmHandler  *app.StartVmHandler
	stopVmHandler   *app.StopVmHandler
}

func New(cfg *config.Config, vmiCli vmi.VirtualMachineInterface) *Handler {
	vmRepo := adapters.NewVmRepository(db.GetDB())

	h := &Handler{
		Engine:          gin.New(),
		createVmHandler: app.NewCreateVmHandler(vmRepo, vmiCli),
		deleteVmHandler: app.NewDeleteVmHandler(vmRepo, vmiCli),
		getVmsHandler:   app.NewGetVmsHandler(vmRepo),
		startVmHandler:  app.NewStartVmHandler(vmRepo, vmiCli),
		stopVmHandler:   app.NewStopVmHandler(vmRepo, vmiCli),
	}

	h.Use(middleware.Recovery())
	h.Use(middleware.StructuredLoggerWithConfig(middleware.LoggerConfig{
		SkipPaths:       cfg.Log.SkipPaths,
		LogRequestBody:  cfg.Log.LogRequestBody,
		LogResponseBody: cfg.Log.LogResponseBody,
		MaxBodySize:     cfg.Log.MaxBodySize,
		SlowThreshold:   cfg.Log.SlowThreshold,
	}))

	o := h.Group("vm")
	o.POST("/", h.CreateVm)
	o.GET("/", h.GetVms)
	o.POST("/:id/delete", h.DeleteVm)
	o.POST("/:id/start", h.StartVm)
	o.POST("/:id/stop", h.StopVm)

	return h
}
