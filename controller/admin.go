package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dao"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
)

type AdminController struct {}

func AdminRegister(group *gin.RouterGroup)  {
	adminLogin:= &AdminController{}
	group.GET("/admin_info",adminLogin.AdminInfo)
	group.POST("/change_pwd",adminLogin.ChangePwd)

}

// AdminInfo godoc
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 获取用户信息
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminlogin *AdminController) AdminInfo(c *gin.Context) {

	sess:=sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)),adminSessionInfo); err != nil {
		middleware.ResponseError(c,2009,err)
		return
	}

	out :=&dto.AdminInfoOutput{
		ID: adminSessionInfo.ID,
		UserName: adminSessionInfo.UserName,
		LoginTime: adminSessionInfo.LoginTime,
		Avatar: "http://www.baidu.com",
		Introduction: "http://www.baidu.com",
		Roles: []string{"admin"},
	}
	middleware.ResponseSuccess(c,out)
}


// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param polygon body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (adminlogin *AdminController) ChangePwd(c *gin.Context) {

	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c,2000,err)
		return
	}

	sess:=sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)),adminSessionInfo); err != nil {
		middleware.ResponseError(c,2009,err)
		return
	}

	adminInfo := &dao.Admin{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	adminInfo,err = adminInfo.Find(c,tx,(&dao.Admin{UserName: adminSessionInfo.UserName}))
	if err !=nil {
		middleware.ResponseError(c,2001,err)
		return
	}
	saltPassword := public.GenSaltPassword(adminInfo.Salt,params.Password)
	adminInfo.Password = saltPassword
	if  err :=adminInfo.Save(c,tx); err !=nil {
		middleware.ResponseError(c,2003,err)
	}
	middleware.ResponseSuccess(c,"")
}