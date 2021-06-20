package controller

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"strings"
)

type ServiceController struct {}

func ServiceRegister(group *gin.RouterGroup)  {
	service := &ServiceController{}
	group.GET("/service_list",service.ServiceList)
	group.POST("/service_delete",service.ServiceDelete)
	group.GET("/service_detail", service.ServiceDetail)
	group.POST("/service_add_http",service.ServiceAddHTTP)
	group.POST("/service_update_http",service.ServiceUpdateHTTP)
}

func (service *ServiceController) ServiceDetail(c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.ServiceDeleteParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, serviceDetail)
}


// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {

	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err!=nil {
		middleware.ResponseError(c,2000,err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	list,total,err := serviceInfo.PageList(c,tx,params)

	if err !=nil {
		middleware.ResponseError(c,2001,err)
	}

	outList := []dto.ServiceListItemOutput{}

	for _,listItem := range list {
		serviceDetail,err := listItem.ServiceDetail(c,tx,&listItem)
		if err !=nil {
			middleware.ResponseError(c,2004,err)
			return
		}

		// 1.http后缀接入 clusterIP+clusterPort+path
		// 2.http域名接入： domain
		// 3.tcp\grpc 接入 clusterIp+servicePort
		serviceAddr := "unknow"
		clusterIp := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

		if serviceDetail.Info.LoadType  == public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType ==public.HTTPPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 1{
			serviceAddr = fmt.Sprintf("%s:%s%s",clusterIp, clusterSSLPort , serviceDetail.HTTPRule.Rule)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType ==public.HTTPPRuleTypePrefixURL  &&
			serviceDetail.HTTPRule.NeedHttps == 0{
			serviceAddr = fmt.Sprintf("%s:%s%s",clusterIp , clusterPort , serviceDetail.HTTPRule.Rule)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType ==public.HTTPRuleTypeDomain{
			serviceAddr = fmt.Sprintf("%s",serviceDetail.HTTPRule.Rule)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeTCP{
			serviceAddr = fmt.Sprintf("%s:%d", clusterIp,serviceDetail.TCPRule.Port)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeGRPC{
			serviceAddr = fmt.Sprintf("%s:%d", clusterIp,serviceDetail.GRPCRule.Port)
		}

		ipList :=serviceDetail.LoadBalance.GetIPListByModel()
		outItem:=dto.ServiceListItemOutput{
			ID: listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps:0,
			Qpd:0,
			TotalNode: len(ipList),
		}

		outList = append(outList,outItem)
	}

	out := &dto.ServiceListOutput{
		List: outList,
		Total:total,
	}

	middleware.ResponseSuccess(c,out)
}


func (service *ServiceController) ServiceDelete(c *gin.Context) {

	params := &dto.ServiceDeleteInput{}
	if err := params.ServiceDeleteParam(c); err!=nil {
		middleware.ResponseError(c,2000,err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}


	serviceInfo := &dao.ServiceInfo{ID:params.ID}
	serviceInfo,err = serviceInfo.Find(c,tx,serviceInfo)

	if err !=nil {
		middleware.ResponseError(c,2001,err)
	}

	serviceInfo.IsDelete = 1
	if err := serviceInfo.Save(c,tx,);err != nil {
		middleware.ResponseError(c,2008,err)
		return
	}

	middleware.ResponseSuccess(c,"success")
}

func (service *ServiceController) ServiceAddHTTP(c *gin.Context) {

	params := &dto.ServiceAddHTTPInput{}
	if err := params.BindValidParam(c); err!=nil {
		middleware.ResponseError(c,2000,err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo,err = serviceInfo.Find(c,tx,serviceInfo)
	if err == nil {
		middleware.ResponseError(c,2002,errors.New("服务已经存在"))
		return
	}
	httpUrl := &dao.HttpRule{RuleType: params.RuleType,Rule: params.Rule}
	if _,err := httpUrl.Find(c,tx,httpUrl);err == nil {
		middleware.ResponseError(c,2003,errors.New("接入域名前缀已经存在"))
		return
	}

	if len(strings.Split(params.IpList,"\n")) != len(strings.Split(params.WeightList,"\n")) {
		middleware.ResponseError(c,2004,errors.New("ip列表和权重不一致"))
		return
	}
	serviceModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := serviceModel.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}
	//serviceModel.ID
	httpRule := &dao.HttpRule{
		ServiceID:      serviceModel.ID,
		RuleType:       params.RuleType,
		Rule:           params.Rule,
		NeedHttps:      params.NeedHttps,
		NeedStripUri:   params.NeedStripUri,
		NeedWebsocket:  params.NeedWebsocket,
		UrlRewrite:     params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         serviceModel.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		ClientIPFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	loadbalance := &dao.LoadBalance{
		ServiceID:              serviceModel.ID,
		RoundType:              params.RoundType,
		IpList:                 params.IpList,
		WeightList:             params.WeightList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout:  params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout:    params.UpstreamIdleTimeout,
		UpstreamMaxIdle:        params.UpstreamMaxIdle,
	}
	if err := loadbalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")

}

func (service *ServiceController) ServiceUpdateHTTP(c *gin.Context) {

	params := &dto.ServiceUpdateHTTPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2001, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, errors.New("服务不存在"))
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(c, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, errors.New("服务不存在"))
		return
	}

	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	httpRule := serviceDetail.HTTPRule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.Rule = params.Rule
	httpRule.NeedStripUri = params.NeedStripUri
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientIPFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	loadbalance := serviceDetail.LoadBalance
	loadbalance.RoundType = params.RoundType
	loadbalance.IpList = params.IpList
	loadbalance.WeightList = params.WeightList
	loadbalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadbalance.UpstreamHeaderTimeout = params.UpstreamHeaderTimeout
	loadbalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadbalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err := loadbalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")

}