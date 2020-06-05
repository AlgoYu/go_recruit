package controllers

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
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
		o := orm.NewOrm()
		account.Picture = DEFAULT_PICTURE_URL
		var hash []byte
		hash,err = bcrypt.GenerateFromPassword([]byte(account.Password),bcrypt.DefaultCost)
		account.Password = string(hash)
		account.CreateDatetime = time.Now()
		account.UpdateDatetime = time.Now()
		var accountId int64
		if accountId, err = o.Insert(&account);err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else{
			accountController.Data["json"] = common.Success(accountId)
		}
	}
	accountController.ServeJSON()
}

// 查询账户
func (accountController *AccountController)SearchAccount()  {
	id, err := accountController.GetUint64("id")
	if err!=nil{
		accountController.Data["json"] = common.Fail(err.Error())
	}else{
		o := orm.NewOrm()
		account := models.Account{Id: uint(id)}
		err = o.Read(&account)
		if err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else {
			accountController.Data["json"] = common.Success(account)
		}
	}
	accountController.ServeJSON()
}

// 模板
func (accountController *AccountController)Template()  {
	accountController.Ctx.WriteString("Hello world!")
}