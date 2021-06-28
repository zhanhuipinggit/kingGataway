package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
)

//匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建reverseproxy
		// 使用reverseproxy.serverHttp(c.request,c.Response)
	}
}
