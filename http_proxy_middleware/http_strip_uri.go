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
func HTTPStripUriMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		if serviceDetail.HTTPRule.RuleType == public.HTTPPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedStripUri == 1 {
			fmt.Println("old",c.Request.URL.Path)
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path,serviceDetail.HTTPRule.Rule,"",1)
			fmt.Println("new",c.Request.URL.Path)
			}

		c.Next()
	}
}
