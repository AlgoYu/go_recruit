package main

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/controllers"
	"cn.anydevelop/go_recruit/filter"
	"cn.anydevelop/go_recruit/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
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

	// 获取百度API Token
	common.GetBaiduToken()
}

func main() {
	defer common.CloseRedis()
	SetFilter()
	SetRouter()
	beego.Run()
}

func SetFilter() {
	// 设置错误处理
	beego.ErrorHandler("500", common.ErrorHandler)
	// 允许跨域请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	// 最后一个参数必须设置为false 不然无法打印数据
	beego.InsertFilter("/*", beego.FinishRouter, filter.FilterLog, false)
	// 管理员Token过滤
	beego.InsertFilter("/admin/*", beego.BeforeRouter, filter.AdminTokenFilter)
}

func SetRouter() {
	// 文件上传接口
	beego.Router("/upload/picture", &controllers.UploadController{}, "post:UploadPicture")
	// 账户接口
	beego.Router("/account/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/account/logout", &controllers.AccountController{}, "delete:Logout")
	beego.Router("/account/add", &controllers.AccountController{}, "post:AddAccount")
	beego.Router("/account/delete", &controllers.AccountController{}, "delete:DeleteAccount")
	beego.Router("/account/modify", &controllers.AccountController{}, "put:ModifyAccount")
	beego.Router("/account/search", &controllers.AccountController{}, "get:SearchAccount")
	beego.Router("/account/all", &controllers.AccountController{}, "get:AllAccount")
	// 管理员接口
	beego.Router("/loginAdmin", &controllers.AdminController{}, "post:LoginAdmin")
	beego.Router("/admin/testLogin", &controllers.AdminController{}, "get:TestLogin")
	beego.Router("/admin/logout", &controllers.AdminController{}, "delete:LogoutAdmin")
	beego.Router("/admin/add", &controllers.AdminController{}, "post:AddAdmin")
	beego.Router("/admin/delete", &controllers.AdminController{}, "delete:DeleteAdmin")
	beego.Router("/admin/search", &controllers.AdminController{}, "get:SearchAdmin")
	beego.Router("/admin/delete", &controllers.AdminController{}, "put:ModifyAdmin")
}
