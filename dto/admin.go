package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

type AdminInfoOutput struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
	Avatar string `json:"avatar"`
	Introduction string `json:"introduction"`
	Roles []string `json:"roles"`
}

type ChangePwdInput struct {
	Password string `form:"password" json:"password" comment:"密码"  validate:"required" example:"123456"`
}

func (param *ChangePwdInput) BindValidParam(c *gin.Context)  error {
	return public.DefaultGetValidParams(c,param)
}