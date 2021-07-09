package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

type AdminLogin struct {}

func AdminLoginRegister(group *gin.RouterGroup)  {
	adminLogin:= &AdminLogin{}
	group.POST("/login",adminLogin.AdminLogin)
	group.GET("/login_out",adminLogin.AdminLoginOut)


}

func (adminlogin *AdminLogin) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err !=nil {
		middleware.ResponseError(c,1001,err)
		return
	}

	admin := dao.Admin{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	admins,err := admin.LoginCheck(c,tx,params)
	if err !=nil {
		middleware.ResponseError(c,2001,err)
		return
	}

	// 设置session
	sessionInfo := &dto.AdminSessionInfo{
		ID: admins.Id,
		UserName: admins.UserName,
		LoginTime: time.Now(),
	}
	sessBts,err := json.Marshal(sessionInfo)
	if err != nil {
		middleware.ResponseError(c,2003,err)
		return
	}

	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey,string(sessBts))
	sess.Save()


	out :=&dto.AdminLoginOutput{Token: admins.UserName}

	middleware.ResponseSuccess(c,out)
}


func (adminlogin *AdminLogin) AdminLoginOut(c *gin.Context) {

	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()

	middleware.ResponseSuccess(c,"")
}

