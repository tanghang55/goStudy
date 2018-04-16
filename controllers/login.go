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
	beego.ReadFromRequest(&l.Controller)
	if l.isPost() {
		flash := beego.NewFlash()
		errorMsg := l.LoginAdmin()
		flash.Error(errorMsg)
		flash.Store(&l.Controller)
		l.redirect(beego.URLFor("LoginController.Login"))
	}
	l.TplName = "login/login.html"
}
