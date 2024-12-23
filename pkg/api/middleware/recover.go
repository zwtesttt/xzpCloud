package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求处理前，执行中间件逻辑
		defer func() {
			if err := recover(); err != nil {
				// 如果发生了 panic，打印堆栈信息
				stackTrace := string(debug.Stack())
				fmt.Printf("Panic: %v\nStackTrace: %s\n", err, stackTrace)

				// 返回错误信息给客户端
				api.RenderInternalServerError(c, fmt.Errorf("Internal Server Error"))
				c.Abort() // 终止后续的处理
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
