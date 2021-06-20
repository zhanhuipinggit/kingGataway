package controller

import (
	"errors"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

type DashboardController struct {}

func DashboardRegister(group *gin.RouterGroup)  {
	dashboard := &DashboardController{}
	group.GET("/dashboard",dashboard.Dashboard)
	group.GET("/flowStat",dashboard.FlowStat)
	group.GET("/service_stat",dashboard.ServiceStat)
}

func (service *DashboardController) ServiceStat(c *gin.Context) {

	tx,err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c,2001,err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	list,err := serviceInfo.GroupByLoadType(c,tx)
	if err !=nil {
		middleware.ResponseError(c,2002,err)
		return
	}

	legend := []string{}
	for index, item := range list{
		name,ok := public.LoadTypeMap[item.LoadType]
		if !ok {
			middleware.ResponseError(c,2003,errors.New("load not fount"))
			return
		}
		list[index].Name=name
		legend = append(legend,name)
	}

	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data: list,

	}

	middleware.ResponseSuccess(c,out)
}




func (admin *DashboardController) FlowStat(c *gin.Context) {

	//今日流量全天小时级访问统计
	todayStat := []int64{}
	for i := 0; i <= time.Now().In(lib.TimeLocation).Hour(); i++ {
		todayStat = append(todayStat, 0)
	}

	//昨日流量全天小时级访问统计
	yesterdayStat := []int64{}
	for i := 0; i <= 23; i++ {
		yesterdayStat = append(yesterdayStat, 0)
	}
	stat := dto.StatisticsOutput{
		Today:     todayStat,
		Yesterday: yesterdayStat,
	}
	middleware.ResponseSuccess(c, stat)
	return
}





func (service *DashboardController) Dashboard(c *gin.Context) {

	tx,err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c,2001,err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	_,serviceNum,err := serviceInfo.PageList(c,tx,&dto.ServiceListInput{PageSize: 1,PageNo: 1})
	if err !=nil {
		middleware.ResponseError(c,2002,err)
		return
	}

	app := &dao.App{}
	_,appNum,err := app.APPList(c,tx,&dto.APPListInput{PageNo: 1,PageSize: 1})
	if err !=nil {
		middleware.ResponseError(c,2002,err)
		return
	}



	out := &dto.PanelGroupData{
		ServiceNum: serviceNum,
		AppNum: appNum,
		TodayRequestNum: 0,
		CurrentQPS: 0,


	}

	middleware.ResponseSuccess(c,out)
}


