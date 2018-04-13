package controllers

import (
	"github.com/astaxie/beego"
	"goStudy/models"
	"goStudy/libs"
	"time"
	"strings"
	"strconv"
)

type BaseController struct {
	beego.Controller
	admin models.Admin
}

func (b *BaseController) getPost(model interface{}) error {
	return b.ParseForm(&model)
}
// 是否POST提交
func (b *BaseController) isPost() bool {
	return b.Ctx.Request.Method == "POST"
}

func (b *BaseController) LoginAdmin() string {
	admin:=b.admin
	errorMsg := ""
	if nil != b.getPost(admin) {
		user, err := admin.GetUser(admin.LoginName)
		if err != nil || user.Password != libs.Md5([]byte(admin.Password+user.Salt)) {
			errorMsg = "帐号或密码错误"
		} else if user.Status == -1 {
			errorMsg = "该帐号已禁用"
		} else {
			user.LastIp = b.getClientIp()
			user.LastLogin = time.Now().Unix()
			user.Update()
			authkey := libs.Md5([]byte(b.getClientIp() + "|" + user.Password + user.Salt))
			b.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
			b.redirect(beego.URLFor("HomeController.Index"))
		}
	}
	return errorMsg
}

//获取用户IP地址
func (b *BaseController) getClientIp() string {
	s := b.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (b *BaseController) redirect(url string) {
	b.Redirect(url, 302)
	b.StopRun()
}