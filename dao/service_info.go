package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

type ServiceInfo struct {
ID        int64       `json:"id" gorm:"primary_key" description:"自增主键"`
LoadType  int    `json:"load_type" gorm:"column:load_type" description:"负载均衡类型 0=http 1=tcp 2=grpc"`
ServiceName    string       `json:"service_name" gorm:"column:service_name" description:"服务名称"`
ServiceDesc    string     `json:"service_desc" gorm:"service_desc:salt" description:"服务描述"`
UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
IsDelete int8 `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *ServiceInfo) GroupByLoadType(c *gin.Context, tx *gorm.DB) ([]dto.DashServiceStatItemOutput, error) {
	list := []dto.DashServiceStatItemOutput{}
	query := tx.SetCtx(public.GetGinTraceContext(c))
	if err := query.Table(t.TableName()).Where("is_delete=0").Select("load_type, count(*) as value").Group("load_type").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo) ServiceDetail (c *gin.Context, tx *gorm.DB,search *ServiceInfo) (*ServiceDetail,error) {
	if search.ServiceName == "" {
		info,err := t.Find(c,tx,search)
		if err != nil {
			return nil,err
		}
		search = info
	}

	httpRule:=&HttpRule{ServiceID: search.ID}
	httpRule,err := httpRule.Find(c,tx,httpRule)
	if err!=nil && err != gorm.ErrRecordNotFound {
		return nil,err
	}
	grpcRule:=&GrpcRule{ServiceID: search.ID}
	grpcRule,err = grpcRule.Find(c,tx,grpcRule)
	if err!=nil && err != gorm.ErrRecordNotFound {
		return nil,err
	}

	tcpRule:=&TcpRule{ServiceID: search.ID}
	tcpRule,err = tcpRule.Find(c,tx,tcpRule)
	if err!=nil && err != gorm.ErrRecordNotFound {
		return nil,err
	}

	accessControl:=&AccessControl{ServiceID: search.ID}
	accessControl,err = accessControl.Find(c,tx,accessControl)
	if err!=nil && err != gorm.ErrRecordNotFound {
		return nil,err
	}

	loadBalance:=&LoadBalance{ServiceID: search.ID}
	loadBalance,err = loadBalance.Find(c,tx,loadBalance)
	if err!=nil && err != gorm.ErrRecordNotFound {
		return nil,err
	}

	detail := &ServiceDetail{
		Info:          search,
		HTTPRule:     httpRule,
		GRPCRule:      grpcRule,
		TCPRule:       tcpRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}

	return detail,nil



}


func (t *ServiceInfo) PageList (c *gin.Context, tx *gorm.DB,params *dto.ServiceListInput) ([]ServiceInfo,int64,error)  {

	total := int64(0)
	list := []ServiceInfo{}
	offset := (params.PageNo-1)*params.PageSize
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query  = query.Table(t.TableName()).Where("is_delete = 0")
	if params.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)","%"+params.Info+"%","%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error; err !=nil && err !=gorm.ErrRecordNotFound {
		return nil,0,err
	}

	query.Count(&total)
	return list,total,nil


}



func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out:=&ServiceInfo{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}