package main

import (
	"errors"
	"os"
	"time"

	"github.com/zwtesttt/xzpCloud/pkg/log"
)

func main() {
	println("=== 日志包演示 ===\n")

	// 演示基础的结构化日志
	println("1. 基础结构化日志：")
	log.StructuredInfo("应用程序启动", 
		log.String("service", "demo"),
		log.String("version", "v1.0.0"),
		log.Int("port", 8080))

	log.StructuredDebug("调试信息", 
		log.String("module", "auth"),
		log.Bool("debug_mode", true))

	log.StructuredWarn("警告信息", 
		log.String("reason", "配置文件未找到"),
		log.String("fallback", "使用默认配置"))

	log.StructuredError("错误信息", 
		log.String("operation", "connect_database"),
		log.ErrorField(errors.New("连接超时")),
		log.Int("retry_count", 3))

	println("\n2. JSON格式输出：")
	// 切换到JSON格式
	log.SetStructuredFormat(true)
	
	log.StructuredInfo("用户登录成功",
		log.String("username", "admin"),
		log.String("ip", "192.168.1.100"),
		log.Time("login_time", time.Now()),
		log.Duration("response_time", 150*time.Millisecond))

	log.StructuredError("API请求失败",
		log.String("method", "POST"),
		log.String("url", "/api/users"),
		log.Int("status_code", 500),
		log.ErrorField(errors.New("内部服务器错误")))

	println("\n3. 日志级别过滤演示：")
	// 设置为ERROR级别
	log.SetStructuredLevel(log.ErrorLevel)
	log.SetStructuredFormat(false) // 切换回文本格式

	println("   设置日志级别为ERROR，以下Info和Warn不会输出：")
	log.StructuredInfo("这条信息不会显示")
	log.StructuredWarn("这条警告不会显示")
	log.StructuredError("只有这条错误会显示", log.String("level", "ERROR"))

	// 恢复日志级别
	log.SetStructuredLevel(log.DebugLevel)
	println("\n   恢复到DEBUG级别，所有日志都会输出：")
	
	log.StructuredDebug("现在可以看到调试信息了")
	log.StructuredInfo("信息级别日志")
	log.StructuredWarn("警告级别日志")
	log.StructuredError("错误级别日志")

	println("\n4. 自定义日志器演示：")
	// 创建自定义日志器
	customLogger := log.NewStructuredLogger(log.InfoLevel, os.Stdout, true)
	
	customLogger.Info("自定义日志器 - JSON格式",
		log.String("logger", "custom"),
		log.Bool("json", true))

	customLogger.Warn("自定义日志器警告",
		log.String("component", "database"),
		log.String("issue", "连接数过多"))

	println("\n=== 演示完成 ===")
}