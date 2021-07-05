package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"strings"
)

//匹配接入方式 基于请求信息
func HTTPWhileListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		iplist := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			iplist = strings.Split(serviceDetail.AccessControl.WhiteList,",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(iplist)>0 {
			if !public.InStringSlice(iplist,c.ClientIP()) {
				middleware.ResponseError(c,3001,errors.New( fmt.Sprintf("%s not in white ip list",c.ClientIP())))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func HTTPBlackMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		serverInterface,ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c,2001,errors.New("service not found"))
			c.Abort()
			return
		}

		serviceDetail := serverInterface.(*dao.ServiceDetail)

		ipBlackList := []string{}
		if serviceDetail.AccessControl.BlackList != "" {
			ipBlackList = strings.Split(serviceDetail.AccessControl.BlackList,",")
		}

		if serviceDetail.AccessControl.OpenAuth == 1 && len(ipBlackList)>0 {
			if public.InStringSlice(ipBlackList,c.ClientIP()) {
				middleware.ResponseError(c,3001,errors.New( fmt.Sprintf("%s not in white ip list",c.ClientIP())))
				c.Abort()
				return
			}
		}
		c.Next()

	}
}

