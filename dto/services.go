package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/public"
)

func (param *ServiceUpdateHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceUpdateHTTPInput struct {
	ID          int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"required,min=1"`                                                     //服务ID
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"test_http_service_indb" validate:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"test_http_service_indb" validate:"required,max=255,min=1"`     //服务描述

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"max=1,min=0"`                             //接入类型
	Rule           string `json:"rule" form:"rule" comment:"接入路径：域名或者前缀" example:"/test_http_service_indb" validate:"required,valid_rule"` //域名或者前缀
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"max=1,min=0"`                        //支持https
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"" validate:"max=1,min=0"`            //启用strip_uri
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"" validate:"max=1,min=0"`          //是否支持websocket
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" example:"" validate:"valid_url_rewrite"`                //url重写功能
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:"valid_header_transfor"`   //header转换

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" example:"" validate:"max=1,min=0"`                     //关键词
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单ip" example:"" validate:""`                               //黑名单ip
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单ip" example:"" validate:""`                               //白名单ip
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流	" example:"" validate:"min=0"` //客户端ip限流
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`          //服务端限流

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"127.0.0.1:80" validate:"required,valid_ipportlist"`                //ip列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"50" validate:"required,valid_weightlist"`             //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" validate:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" validate:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" validate:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`                     //最大空闲链接数
}

func (param *ServiceAddHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceAddHTTPInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`     //服务描述

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"max=1,min=0"`                           //接入类型
	Rule           string `json:"rule" form:"rule" comment:"接入路径：域名或者前缀" example:"" validate:"required,valid_rule"`                      //域名或者前缀
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"max=1,min=0"`                      //支持https
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri" example:"" validate:"max=1,min=0"`          //启用strip_uri
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"" validate:"max=1,min=0"`        //是否支持websocket
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" example:"" validate:"valid_url_rewrite"`              //url重写功能
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header转换" example:"" validate:"valid_header_transfor"` //header转换

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限" example:"" validate:"max=1,min=0"`                     //关键词
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单ip" example:"" validate:""`                               //黑名单ip
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单ip" example:"" validate:""`                               //白名单ip
	ClientipFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流	" example:"" validate:"min=0"` //客户端ip限流
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`          //服务端限流

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ipportlist"`                            //ip列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weightlist"`               //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" validate:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" validate:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" validate:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`                     //最大空闲链接数
}

type ServiceListInput struct {
	Info string `form:"info" json:"info" comment:"关键词"  validate:"" example:""`
	PageNo int `form:"page_no" json:"page_no" comment:"页数"  validate:"required" example:"1"`
	PageSize int `form:"page_size" json:"page_size" comment:"每页数"  validate:"required" example:"20"`
}

type ServiceDeleteInput struct {
	ID int64 `form:"id" json:"id" comment:"id"  validate:"required" example:"56"`
}
func (param *ServiceDeleteInput) ServiceDeleteParam(c *gin.Context)  error {
	return public.DefaultGetValidParams(c,param)
}

func (param *ServiceListInput) BindValidParam(c *gin.Context)  error {
	return public.DefaultGetValidParams(c,param)
}

type ServiceListItemOutput struct {
	ID int64 `json:"id" form:"id"`
	ServiceName string `json:"ServiceName" form:"ServiceName"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType int `json:"load_type" form:"load_type"`
	ServiceAddr string `json:"service_addr" form:"service_addr"`
	Qps int64 `json:"qps" form:"qps"`
	Qpd int64 `json:"qpd" form:"qpd"`
	TotalNode int `json:"total_node" form:"total_node"`
}

type ServiceListOutput struct {
	Total int64 `form:"total" json:"total" comment:"总数"  validate:"" example:""`
	List []ServiceListItemOutput `form:"list" json:"list" comment:"数据"  validate:"" example:""`
}