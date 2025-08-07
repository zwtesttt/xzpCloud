package main

import (
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/pkg/api/middleware"
	"github.com/zwtesttt/xzpCloud/pkg/log"
)

// 这个文件展示如何在现有的handler中集成新的日志中间件
// 以VM服务为例

// 示例：如何更新 internal/vm/api/handler/handler.go
func ExampleVMHandlerIntegration() {
	fmt.Println("=== VM Handler 集成示例 ===")
	
	// 创建自定义日志配置
	logConfig := middleware.LoggerConfig{
		SkipPaths:       []string{"/health", "/metrics"},
		LogRequestBody:  true,
		LogResponseBody: true,
		MaxBodySize:     2048, // 2KB
		SlowThreshold:   100 * time.Millisecond, // VM操作可能较慢，设置更低的阈值
	}

	// 模拟VM Handler的创建过程
	h := gin.New()
	
	// 1. 添加基础中间件
	h.Use(middleware.Recovery())
	h.Use(middleware.RequestIDMiddleware())
	h.Use(middleware.StructuredLoggerWithConfig(logConfig))

	// 2. 设置路由组
	vm := h.Group("vm")
	
	// 模拟路由（实际应用中这些会调用真实的handler方法）
	vm.POST("/", func(c *gin.Context) {
		// 在handler内部也可以使用结构化日志
		log.StructuredInfo("开始创建虚拟机",
			log.String("request_id", c.GetString("request_id")),
			log.String("user_id", c.GetHeader("User-ID")),
		)
		
		c.JSON(201, gin.H{"message": "VM创建成功", "vm_id": "vm-123"})
	})
	
	vm.GET("/", func(c *gin.Context) {
		log.StructuredDebug("查询虚拟机列表",
			log.String("request_id", c.GetString("request_id")),
		)
		
		c.JSON(200, gin.H{"vms": []string{"vm-1", "vm-2"}})
	})

	fmt.Println("VM Handler 配置完成")
}

// 示例：如何更新 internal/user/api/handler/handler.go
func ExampleUserHandlerIntegration() {
	fmt.Println("=== User Handler 集成示例 ===")
	
	// 用户服务的日志配置
	logConfig := middleware.LoggerConfig{
		SkipPaths:       []string{"/health"},
		LogRequestBody:  false, // 用户登录等敏感操作不记录请求体
		LogResponseBody: false, // 不记录包含用户信息的响应体
		MaxBodySize:     1024,
		SlowThreshold:   50 * time.Millisecond,
	}

	h := gin.New()
	h.Use(middleware.Recovery())
	h.Use(middleware.RequestIDMiddleware())
	h.Use(middleware.StructuredLoggerWithConfig(logConfig))

	user := h.Group("/user")
	user.POST("/login", func(c *gin.Context) {
		// 在敏感操作中使用日志
		log.StructuredInfo("用户登录尝试",
			log.String("request_id", c.GetString("request_id")),
			log.String("ip", c.ClientIP()),
			log.String("user_agent", c.GetHeader("User-Agent")),
		)
		
		c.JSON(200, gin.H{"message": "登录成功", "token": "***"})
	})

	fmt.Println("User Handler 配置完成")
}

// 示例：现有代码的改造指南
func ExampleMigrationGuide() {
	fmt.Println("\n=== 现有代码改造指南 ===")
	
	fmt.Println(`
1. 更新 internal/vm/api/handler/handler.go:

   原代码:
   func New(vmiCli vmi.VirtualMachineInterface) *Handler {
       vmRepo := adapters.NewVmRepository(db.GetDB())
       h := &Handler{
           Engine: gin.New(),
           // ... 其他字段
       }
       
       o := h.Group("vm")
       // ... 路由设置
       return h
   }

   更新后:
   func New(vmiCli vmi.VirtualMachineInterface) *Handler {
       vmRepo := adapters.NewVmRepository(db.GetDB())
       h := &Handler{
           Engine: gin.New(),
           // ... 其他字段
       }
       
       // 添加中间件
       h.Use(middleware.Recovery())
       h.Use(middleware.RequestIDMiddleware())
       h.Use(middleware.StructuredLogger())
       
       o := h.Group("vm")
       // ... 路由设置
       return h
   }

2. 在具体的handler方法中使用结构化日志:

   原代码:
   func (h *Handler) CreateVm(c *gin.Context) {
       fmt.Println("Creating VM...")  // 替换这种日志
       // ... 业务逻辑
   }

   更新后:
   func (h *Handler) CreateVm(c *gin.Context) {
       log.StructuredInfo("开始创建虚拟机",
           log.String("request_id", c.GetString("request_id")),
           log.String("user_id", c.GetHeader("User-ID")),
       )
       // ... 业务逻辑
   }

3. 错误处理的日志记录:

   原代码:
   if err != nil {
       fmt.Printf("Error: %v", err)  // 替换这种错误日志
       return
   }

   更新后:
   if err != nil {
       log.StructuredError("虚拟机创建失败",
           log.String("request_id", c.GetString("request_id")),
           log.ErrorField(err),
           log.String("operation", "create_vm"),
       )
       return
   }
	`)
}

func main() {
	// 设置日志格式
	log.SetStructuredFormat(false)
	log.SetStructuredLevel(log.InfoLevel)

	ExampleVMHandlerIntegration()
	ExampleUserHandlerIntegration()
	ExampleMigrationGuide()
}