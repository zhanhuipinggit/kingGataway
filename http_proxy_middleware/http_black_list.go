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
func HTTPBlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		whileList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			whileList = strings.Split(serviceDetail.AccessControl.BlackList,",")
		}

		blackList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			blackList = strings.Split(serviceDetail.AccessControl.WhiteList,",")
		}

		if serviceDetail.AccessControl.OpenAuth == 1 && len(whileList)==0  &&
			len(blackList) >0 {
			if !public.InStringSlice(blackList,c.ClientIP()) {
				middleware.ResponseError(c,3001,errors.New( fmt.Sprintf("%s not in blackList ip list",c.ClientIP())))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
