package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (l *LoginController) Login() {
	if l.admin.Id > 0 {
		l.redirect(beego.URLFor("HomeController.Index"))
	}
	if l.isPost() {
		errorMsg:=l.LoginAdmin()
		flash := beego.NewFlash()
		flash.Error(errorMsg)
		flash.Store(&l.Controller)
		l.redirect(beego.URLFor("LoginController.LoginIn"))
	}
	l.TplName = "login/login.html"
}
