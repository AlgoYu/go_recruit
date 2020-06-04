package main

import (
	"cn.anydevelop/go_recruit/controllers"
	"cn.anydevelop/go_recruit/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:xy942698.@tcp(127.0.0.1:3306)/recruit?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.Account))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main()  {
	SetRouter()
	beego.Run()
}

func SetRouter()  {
	beego.Router("/account/add",&controllers.AccountController{},"post:AddAccount")
}