package dao

import (
	"errors"
	"fmt"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/zhanhuipinggit/kingGataway/dto"
	"github.com/zhanhuipinggit/kingGataway/public"
	"time"
)

type Admin struct {
Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
UserName  string    `json:"user_name" gorm:"column:user_name" description:"用户名"`
Password    string       `json:"password" gorm:"column:password" description:"用户名"`
Salt    string     `json:"salt" gorm:"column:salt" description:"盐"`
UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
IsDelete int8 `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *Admin) TableName() string {
	return "gateway_admin"
}

func (t *Admin) LoginCheck (c *gin.Context,db *gorm.DB,input *dto.AdminLoginInput) (*Admin,error)  {

	fmt.Println(input.UserName)
	adminInfo,err :=t.Find(c,db,(&Admin{UserName: input.UserName,IsDelete: 0}))
	if err !=nil {
		return nil,errors.New("用户信息不存在")
	}

	saltPassword := public.GenSaltPassword(adminInfo.Salt,input.Password)
	if saltPassword != adminInfo.Password {
		return nil,errors.New("密码错误，请重新输入")
	}
	
	return adminInfo,nil
}


func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out:=&Admin{}
	fmt.Printf("%v",search)
	print(tx.Debug())
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}