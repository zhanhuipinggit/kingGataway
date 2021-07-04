package http_proxy_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"strings"
)

//匹配接入方式 基于请求信息
func HTTPHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		for _, item := range strings.Split(serviceDetail.HTTPRule.HeaderTransfor,",") {
			
			items :=strings.Split(item," ")
			if len(items) !=3 {
				continue
			}

			if items[0]=="add" || items[0] == "edit" {
				c.Request.Header.Set(items[1],items[2])
			}

			if items[0] == "del" {
				c.Request.Header.Del(items[1])
			}

			
			
			
		}
		c.Next()

	}
}
