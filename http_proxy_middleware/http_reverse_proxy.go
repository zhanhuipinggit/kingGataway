package http_proxy_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/reverse_proxy"
)

//匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}


		serviceDetail := serverInterface.(*dao.ServiceDetail)

		lb,err :=dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)

		if err !=nil {
			middleware.ResponseError(c,2002,errors.New("service ont found"))
			c.Abort()
			return
		}

		trans,err :=dao.TransportHandler.GetTrans(serviceDetail)
		if err != nil {
			middleware.ResponseError(c,2003,err)
			c.Abort()
			return
		}

		proxy :=reverse_proxy.NewLoadBalanceReverseProxy(c,lb,trans)
		proxy.ServeHTTP(c.Writer,c.Request)

	}
}
