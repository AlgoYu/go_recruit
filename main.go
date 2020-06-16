package main

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/controllers"
	"cn.anydevelop/go_recruit/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_connect"), 30)
	//orm.SetMaxIdleConns("default",5)
	//orm.SetMaxOpenConns("default",20)
	//// 设置为 UTC 时间
	//orm.DefaultTimeLoc = time.UTC
	//orm.Debug = true
	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)

	// register model
	orm.RegisterModel(new(models.Account))
	orm.RegisterModel(new(models.Admin))

	// create table
	//orm.RunSyncdb("default", false, true)

	// connect to redis
	common.ConnectRedis()

	common.GetBaiduToken()
}

func main() {
	defer common.CloseRedis()
	SetRouter()
	beego.Run()
}

func SetRouter() {
	beego.Router("/account/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/account/logout", &controllers.AccountController{}, "delete:Logout")
	beego.Router("/account/add", &controllers.AccountController{}, "post:AddAccount")
	beego.Router("/account/delete", &controllers.AccountController{}, "delete:DeleteAccount")
	beego.Router("/account/modify", &controllers.AccountController{}, "put:ModifyAccount")
	beego.Router("/account/search", &controllers.AccountController{}, "get:SearchAccount")
	beego.Router("/account/all", &controllers.AccountController{}, "get:AllAccount")
	beego.Router("/admin/login", &controllers.AccountController{}, "post:LoginAdmin")
	beego.Router("/admin/logout", &controllers.AccountController{}, "delete:LogoutAdmin")
	beego.Router("/admin/add", &controllers.AccountController{}, "post:AddAdmin")
	beego.Router("/admin/delete", &controllers.AccountController{}, "delete:DeleteAdmin")
}
