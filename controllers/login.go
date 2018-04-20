package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	BaseController
}

func (l *LoginController) Login() {
	if l.userId > 0 {
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
	fmt.Println(2)

	l.TplName = "login/login.html"
}
