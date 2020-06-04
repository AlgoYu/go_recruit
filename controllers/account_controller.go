package controllers

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
	"time"
)

const DEFAULT_PICTURE_URL = "http://img.duoziwang.com/2019/05/08050202206333.jpg"

// 账号控制器
type AccountController struct {
	beego.Controller
}

// 增加用户
func (accountController *AccountController)AddAccount()  {
	var account models.Account
	if err := json.Unmarshal(accountController.Ctx.Input.RequestBody,&account);err!=nil{
		accountController.Data["json"] = common.Fail(err.Error())
	}else{
		orm := orm.NewOrm()
		account.Id = uuid.NewV4().String()
		account.Picture = DEFAULT_PICTURE_URL
		account.CreateDatetime = time.Now()
		account.UpdateDatetime = time.Now()
		if insert, err := orm.Insert(&account);err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else{
			accountController.Data["json"] = common.Success(insert)
		}
	}
	accountController.ServeJSON()
}

// 模板
func (accountController *AccountController)Template()  {
	accountController.Ctx.WriteString("Hello world!")
}