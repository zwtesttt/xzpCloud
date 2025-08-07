package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/pkg/api/middleware"
	"github.com/zwtesttt/xzpCloud/pkg/log"
)

func main() {
	// 设置日志格式
	log.SetStructuredFormat(false) // 使用文本格式，方便查看
	log.SetStructuredLevel(log.DebugLevel)

	// 创建gin引擎
	r := gin.New()

	// 添加中间件
	r.Use(middleware.Recovery())           // 恢复中间件
	r.Use(middleware.RequestIDMiddleware()) // 请求ID中间件
	r.Use(middleware.StructuredLogger())   // 结构化日志中间件

	// 定义路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    "2024-01-01 12:00:00",
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	})

	r.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
	})

	r.GET("/slow", func(c *gin.Context) {
		// 模拟慢请求
		time.Sleep(300 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{
			"message": "This was a slow request",
		})
	})

	println("服务器启动在 :8080")
	println("测试URL:")
	println("  GET  http://localhost:8080/ping")
	println("  POST http://localhost:8080/user")
	println("  GET  http://localhost:8080/error")
	println("  GET  http://localhost:8080/slow")
	
	r.Run(":8080")
}

/*
使用示例：

1. 启动服务器：
   go run examples/middleware_usage.go

2. 测试请求：
   curl http://localhost:8080/ping
   curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"张三","email":"zhangsan@example.com"}'
   curl http://localhost:8080/error
   curl http://localhost:8080/slow

3. 查看日志输出，你会看到：
   - 请求开始日志
   - 请求体日志（POST请求）
   - 请求完成日志
   - 响应体日志
   - 错误日志（如果有）
   - 慢请求警告（如果超过阈值）
*/