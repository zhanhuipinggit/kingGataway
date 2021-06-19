package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/public"
)

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