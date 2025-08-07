package config

import (
	"github.com/gin-gonic/gin"
)

// SetupGinMode 设置Gin模式
func SetupGinMode(mode string) {

	switch {
	case mode == "debug":
		gin.SetMode(gin.DebugMode)
	case mode == "test":
		gin.SetMode(gin.TestMode)
	case mode == "release" || mode == "production" || mode == "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		// 默认使用release模式，关闭调试日志
		gin.SetMode(gin.ReleaseMode)
	}
}

// GetGinMode 获取当前Gin模式
func GetGinMode() string {
	return gin.Mode()
}

// IsDebugMode 检查是否为调试模式
func IsDebugMode() bool {
	return gin.Mode() == gin.DebugMode
}

// IsReleaseMode 检查是否为发布模式
func IsReleaseMode() bool {
	return gin.Mode() == gin.ReleaseMode
}
