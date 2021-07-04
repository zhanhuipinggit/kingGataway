package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

//匹配接入方式 基于请求信息
func HTTPFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface,ok:= c.Get("service")

		if !ok {
			middleware.ResponseError(c,2001,errors.New("service ont found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		totalCount,err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil{
			middleware.ResponseError(c,4001,err)
			c.Abort()
			return
		}
		totalCount.Increase()
		dayCount,_:= totalCount.GetDayData(time.Now())
		fmt.Printf("totalCount qps:%v,dayCount:%v",totalCount.QPS,dayCount)

		serviceCounter,err := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix+serviceDetail.Info.ServiceName)
		if err != nil{
			middleware.ResponseError(c,4001,err)
			c.Abort()
			return
		}
		serviceCounter.Increase()

		dayServiceCount,_:= serviceCounter.GetDayData(time.Now())
		fmt.Printf("totalCount qps:%v,dayCount:%v",serviceCounter.QPS,dayServiceCount)
		c.Next()
	}
}
