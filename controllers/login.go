package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	BaseController
}

func (l *LoginController) Login() {
	if l.admin.Id > 0 {
		l.Redirect(beego.URLFor("HomeController.Index"),303)
	}

	l.TplName = "login/login.html"
}
