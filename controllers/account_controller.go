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

// 删除账户
func (accountController *AccountController)DeleteAccount()  {
	id, err := accountController.GetUint64("id")
	if err!=nil{
		accountController.Data["json"] = common.Fail(err.Error())
	}else{
		o := orm.NewOrm()
		var num int64
		if num, err = o.Delete(&models.Account{Id: uint(id)});err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else{
			accountController.Data["json"] = common.Success(num)
		}
	}
	accountController.ServeJSON()
}

// 修改用户
func (accountController *AccountController)ModifyAccount()  {
	var account models.Account
	if err := json.Unmarshal(accountController.Ctx.Input.RequestBody,&account);err!=nil{
		accountController.Data["json"] = common.Fail(err.Error())
	}else{
		o := orm.NewOrm()
		var num int64
		if num, err = o.Update(&account);err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else{
			accountController.Data["json"] = common.Success(num)
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
		if err = o.Read(&account);err!=nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else {
			accountController.Data["json"] = common.Success(account)
		}
	}
	accountController.ServeJSON()
}

// 所有账户
func (accountController *AccountController)AllAccount()  {
	keyWord := accountController.GetString("keyWord")
	currentPage,_ := accountController.GetInt("currentPage")
	pageSize,_ := accountController.GetInt("pageSize")

	o := orm.NewOrm()
	table := o.QueryTable("account")
	table.Filter("name__contains",keyWord)
	table.Filter("email__contains",keyWord)

	var pagesResult common.PageResult
	pagesResult.Total,_ = table.Count()
	pagesResult.PageSize = pageSize
	table.Limit(pageSize,(currentPage-1)*pageSize)

	var accounts []models.Account
	table.All(&accounts)
	pagesResult.Data = accounts
	accountController.Data["json"] = common.Success(pagesResult)
	accountController.ServeJSON()
}