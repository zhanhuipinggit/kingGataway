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
	group.GET("/panel_group_data",dashboard.Dashboard)
	group.GET("/flow_stat",dashboard.FlowStat)
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

	counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	todayList := []int64{}
	currentTime := time.Now()
	for i := 0; i <= currentTime.Hour(); i++ {
		dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourData, _ := counter.GetHourData(dateTime)
		todayList = append(todayList, hourData)
	}

	yesterdayList := []int64{}
	yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourData, _ := counter.GetHourData(dateTime)
		yesterdayList = append(yesterdayList, hourData)
	}
	middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}





func (service *DashboardController) Dashboard(c *gin.Context) {

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(c, tx, &dto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	app := &dao.App{}
	_, appNum, err := app.APPList(c, tx, &dto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}



	out := &dto.PanelGroupData{
		ServiceNum: serviceNum,
		AppNum: appNum,
		TodayRequestNum: counter.TotalCount,
		CurrentQPS:      counter.QPS,


	}

	middleware.ResponseSuccess(c,out)
}


