package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
)

type ServiceController struct {}

func ServiceRegister(group *gin.RouterGroup)  {
	service := &ServiceController{}
	group.GET("/service_list",service.ServiceList)
	group.POST("/service_delete",service.ServiceDelete)
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
			serviceDetail.HTTPPRule.RuleType ==public.HTTPPRuleTypePrefixURL &&
			serviceDetail.HTTPPRule.NeedHttps == 1{
			serviceAddr = fmt.Sprintf("%s:%s%s",clusterIp, clusterSSLPort , serviceDetail.HTTPPRule.Rule)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeHTTP &&
			serviceDetail.HTTPPRule.RuleType ==public.HTTPPRuleTypePrefixURL  &&
			serviceDetail.HTTPPRule.NeedHttps == 0{
			serviceAddr = fmt.Sprintf("%s:%s%s",clusterIp , clusterPort , serviceDetail.HTTPPRule.Rule)
		}

		if serviceDetail.Info.LoadType  == public.LoadTypeHTTP &&
			serviceDetail.HTTPPRule.RuleType ==public.HTTPRuleTypeDomain{
			serviceAddr = fmt.Sprintf("%s",serviceDetail.HTTPPRule.Rule)
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