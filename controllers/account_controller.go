package controllers

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 默认头像
const (
	DEFAULT_PICTURE_URL = "http://img.duoziwang.com/2019/05/08050202206333.jpg"
	MINUTE = 60
)

// 账号控制器
type AccountController struct {
	beego.Controller
}

// 登陆用户
func (accountController *AccountController)Login()  {
	var account models.Account
	if err := json.Unmarshal(accountController.Ctx.Input.RequestBody,&account);err!=nil{
		accountController.Data["json"] = common.Fail(err.Error())
	}else{
		o := orm.NewOrm()
		table := o.QueryTable("account")
		source := models.Account{}
		err = table.Filter("contact",account.Contact).One(&source)
		if err != nil{
			accountController.Data["json"] = common.Fail(err.Error())
		}else{
			err = bcrypt.CompareHashAndPassword([]byte(source.Password),[]byte(account.Password))
			if err!=nil{
				accountController.Data["json"] = common.Fail(err.Error())
			}else{
				token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
					"accountContact": source.Contact,
					"datetime": time.Now(),
				})
				signedString, sigErr := token.SignedString([]byte(beego.AppConfig.String("token_secret_key")))
				if sigErr!=nil{
					beego.Error(sigErr)
				}
				beego.Debug(signedString)
				common.HashPut(source.Contact,"accountId",source.Id)
				common.HashPut(source.Contact,"accountName",source.Name)
				common.HashPut(source.Contact,"accountEmail",source.Email)
				common.Expire(source.Contact,30*MINUTE)
				accountController.Data["json"] = common.Success(signedString)
			}
		}
	}
	accountController.ServeJSON()
}

// 注销登陆
func (accountController *AccountController)Logout()  {
	tokenStr := accountController.Ctx.Request.Header.Get("Token")
	beego.Debug(tokenStr)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("Unexpected signing method: %v\", token.Header[\"alg\"])")
		}
		return []byte(beego.AppConfig.String("token_secret_key")), nil
	})
	if err!=nil{
		beego.Debug(err)
		accountController.Data["json"] = common.Fail(err.Error())
	}else {
		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			beego.Debug(claim["accountContact"].(string))
			beego.Debug(claim["datetime"].(string))
			common.Delete(claim["accountContact"].(string))
			accountController.Data["json"] = common.Success(true)
		}
	}
	accountController.ServeJSON()
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