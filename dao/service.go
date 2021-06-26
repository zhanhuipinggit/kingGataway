package dao

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/public"
	"net/http/httptest"
	"strings"
	"sync"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http_rule" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp_rule" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc_rule" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
}

var ServiceManagerHandler *ServiceManager

func init()  {
	ServiceManagerHandler = NewServiceManager()
}



type ServiceManager struct {
	ServiceMap map[string]*ServiceDetail
	ServiceSlice []*ServiceDetail
	Locker sync.RWMutex
	init sync.Once
	err error
}

func NewServiceManager() *ServiceManager  {
	return &ServiceManager{
		ServiceMap: map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker: sync.RWMutex{},
		init:sync.Once{},
	}
}

func (s *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail,error) {

	// 1.前缀匹配 /abc ==> serviceSlice.rule
	// 2.域名匹配 www.test.com ==> serviceSlice.rule

	// host c.Request.Host
	// path c.Request.URL.path

	host := c.Request.Host
	host = host[0:strings.Index(host,":")]
	fmt.Println("host",host)
	path := c.Request.URL.Path
	fmt.Println("path",path)
	for _,serviceItem := range s.ServiceSlice {
		if serviceItem.Info.LoadType != public.LoadTypeHTTP {
			continue
		}

		if serviceItem.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			if serviceItem.HTTPRule.Rule == host {
				return serviceItem,nil
			}
		}

		if serviceItem.HTTPRule.RuleType == public.HTTPPRuleTypePrefixURL {
			if strings.HasPrefix(path,serviceItem.HTTPRule.Rule) {
				return serviceItem,nil
			}
		}
	}

	return nil,errors.New("is not Server")
}


func (s *ServiceManager) LoadOnce () error {
	s.init.Do(func() {

		c ,_:=gin.CreateTestContext(httptest.NewRecorder())
		tx, err := lib.GetGormPool("default")
		if err != nil {
			s.err = err
			return
		}
		serviceInfo := &ServiceInfo{}
		params := &dto.ServiceListInput{PageNo: 1,PageSize: 999999}
		list,_,err := serviceInfo.PageList(c,tx,params)

		if err !=nil {
			s.err = err
			return
		}
		s.Locker.Lock()
		defer s.Locker.Unlock()
		for _,listItem := range list {
			tmpItem := listItem
			serviceDetail,err := tmpItem.ServiceDetail(c,tx,&tmpItem)
			if err != nil {
				s.err = err
				return
			}
			s.ServiceMap[tmpItem.ServiceName] = serviceDetail
			s.ServiceSlice = append(s.ServiceSlice,serviceDetail)
		}

	})
	return s.err
}


