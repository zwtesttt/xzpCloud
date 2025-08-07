package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/pkg/log"
)

// responseWriter 包装gin的ResponseWriter以捕获响应内容
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// LoggerConfig 日志中间件配置
type LoggerConfig struct {
	// SkipPaths 跳过记录的路径
	SkipPaths []string
	// LogRequestBody 是否记录请求体
	LogRequestBody bool
	// LogResponseBody 是否记录响应体
	LogResponseBody bool
	// MaxBodySize 最大记录的请求/响应体大小（字节）
	MaxBodySize int64
	// SlowThreshold 慢请求阈值
	SlowThreshold time.Duration
}

// DefaultLoggerConfig 默认日志配置
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		SkipPaths:       []string{"/health", "/metrics", "/favicon.ico"},
		LogRequestBody:  true,
		LogResponseBody: true,
		MaxBodySize:     1024 * 4, // 4KB
		SlowThreshold:   200 * time.Millisecond,
	}
}

// Logger 创建日志中间件
func Logger() gin.HandlerFunc {
	return LoggerWithConfig(DefaultLoggerConfig())
}

// LoggerWithConfig 使用自定义配置创建日志中间件
func LoggerWithConfig(config LoggerConfig) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 使用我们自己的日志包而不是gin默认的日志格式
		return ""
	})
}

// StructuredLogger 结构化日志中间件
func StructuredLogger() gin.HandlerFunc {
	return StructuredLoggerWithConfig(DefaultLoggerConfig())
}

// StructuredLoggerWithConfig 使用配置的结构化日志中间件
func StructuredLoggerWithConfig(config LoggerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否需要跳过
		for _, path := range config.SkipPaths {
			if c.Request.URL.Path == path {
				c.Next()
				return
			}
		}

		startTime := time.Now()

		// 读取请求体
		var requestBody string
		if config.LogRequestBody && c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(io.LimitReader(c.Request.Body, config.MaxBodySize))
			if err == nil {
				requestBody = string(bodyBytes)
				// 重新设置请求体供后续处理使用
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 包装ResponseWriter以捕获响应
		var responseBody string
		if config.LogResponseBody {
			w := &responseWriter{
				ResponseWriter: c.Writer,
				body:           bytes.NewBufferString(""),
			}
			c.Writer = w
			defer func() {
				if w.body.Len() <= int(config.MaxBodySize) {
					responseBody = w.body.String()
				} else {
					responseBody = w.body.String()[:config.MaxBodySize] + "...(truncated)"
				}
			}()
		}

		// 记录请求开始
		log.StructuredInfo("HTTP请求开始",
			log.String("method", c.Request.Method),
			log.String("path", c.Request.URL.Path),
			log.String("query", c.Request.URL.RawQuery),
			log.String("ip", c.ClientIP()),
			log.String("user_agent", c.Request.UserAgent()),
		)

		// 如果需要记录请求体且不为空
		if config.LogRequestBody && requestBody != "" && !isHealthCheck(c.Request.URL.Path) {
			log.StructuredDebug("HTTP请求体",
				log.String("method", c.Request.Method),
				log.String("path", c.Request.URL.Path),
				log.String("body", requestBody),
			)
		}

		// 处理请求
		c.Next()

		// 计算处理时间
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		// 构建基础字段
		fields := []log.Field{
			log.String("method", c.Request.Method),
			log.String("path", c.Request.URL.Path),
			log.String("query", c.Request.URL.RawQuery),
			log.String("ip", c.ClientIP()),
			log.Int("status_code", statusCode),
			log.Duration("duration", duration),
		}

		// 添加响应体（如果需要且不为空）
		if config.LogResponseBody && responseBody != "" && !isHealthCheck(c.Request.URL.Path) {
			fields = append(fields, log.String("response_body", responseBody))
		}

		// 根据状态码和处理时间选择日志级别
		var message string
		var logLevel log.Level

		if statusCode >= 500 {
			message = "HTTP请求完成 - 服务器错误"
			logLevel = log.ErrorLevel
		} else if statusCode >= 400 {
			message = "HTTP请求完成 - 客户端错误"
			logLevel = log.WarnLevel
		} else if duration > config.SlowThreshold {
			message = "HTTP请求完成 - 慢请求"
			logLevel = log.WarnLevel
		} else {
			message = "HTTP请求完成"
			logLevel = log.InfoLevel
		}

		// 记录日志
		switch logLevel {
		case log.ErrorLevel:
			log.StructuredError(message, fields...)
		case log.WarnLevel:
			log.StructuredWarn(message, fields...)
		default:
			log.StructuredInfo(message, fields...)
		}

		// 如果有错误，记录错误详情
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.StructuredError("HTTP请求处理错误",
					log.String("method", c.Request.Method),
					log.String("path", c.Request.URL.Path),
					log.ErrorField(err.Err),
					log.Int("error_type", int(err.Type)),
				)
			}
		}
	}
}

// isHealthCheck 判断是否为健康检查请求
func isHealthCheck(path string) bool {
	healthPaths := []string{"/health", "/ping", "/status", "/metrics"}
	for _, hp := range healthPaths {
		if strings.Contains(path, hp) {
			return true
		}
	}
	return false
}

// AccessLogger 访问日志中间件（简化版）
func AccessLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 使用结构化日志记录访问日志
		log.StructuredInfo("访问日志",
			log.String("method", param.Method),
			log.String("path", param.Path),
			log.String("ip", param.ClientIP),
			log.Int("status", param.StatusCode),
			log.Duration("latency", param.Latency),
		)
		return ""
	})
}
