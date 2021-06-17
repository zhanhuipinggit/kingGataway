package controller

import (
	"encoding/json"
	"fmt"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/middleware"
	"github.com/zhanhuipinggit/kingGataway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {}

func AdminRegister(group *gin.RouterGroup)  {
	adminLogin:= &AdminController{}
	group.GET("/admin_info",adminLogin.AdminInfo)

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
