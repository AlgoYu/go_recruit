package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

// 登陆用户
func (this *UserController)Login()  {
	this.Ctx.WriteString("Hello world!")
}