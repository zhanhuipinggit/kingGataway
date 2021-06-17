package dto

import (
	"github.com/zhanhuipinggit/kingGataway/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminLoginInput struct {
	UserName string `form:"user_name" json:"user_name" comment:"用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"  validate:"required" example:"123456"`
}

type AdminSessionInfo struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginOutput struct {
	Token string `json:"token" comment:"用户名"  form:"token" comment:"token" example:"admin"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context)  error {
	return public.DefaultGetValidParams(c,param)
}