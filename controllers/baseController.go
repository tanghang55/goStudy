package controllers

import (
	"github.com/astaxie/beego"
	"goStudy/models"
	"goStudy/libs"
	"time"
	"strings"
	"strconv"
	"fmt"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	admin          *models.Admin
	pageSize       int
	controllerName string
	actionName     string
	userId         int
	userName       string
	allowUrl       string
}

//准备
func (b *BaseController) Prepare() {
	b.pageSize = 20
	controllerName, actionName := b.GetControllerAndAction()                     //获得控制名和方法名
	b.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10]) //控制名
	b.actionName = strings.ToLower(actionName)                                   //方法名
	b.Data["version"] = beego.AppConfig.String("version")
	b.Data["siteName"] = beego.AppConfig.String("site.name")
	b.Data["curRoute"] = b.controllerName + "." + b.actionName
	b.Data["curController"] = b.controllerName
	b.Data["curAction"] = b.actionName
	if (strings.Compare(b.controllerName, "apidoc")) != 0 {
		b.auth()
	}
	b.Data["loginUserId"] = b.userId
	b.Data["loginUserName"] = b.userName

}

//登录权限验证
func (b *BaseController) auth() {
	arr := strings.Split(b.Ctx.GetCookie("auth"), "|") //切割cookie值
	b.userId = 0                                       //初始化id
	if len(arr) == 2 {
		uid, pwd := arr[0], arr[1]
		userId, _ := strconv.Atoi(uid) //字符串转数字
		if userId > 0 {
			admin := models.Admin{}
			user, err := admin.FindIdUser(userId)
			if err == nil && pwd == libs.Md5([]byte(b.getClientIp()+"|"+user.Password+user.Salt)) {
				b.userId = user.Id
				b.userName = user.LoginName
				b.admin = user
				//b.AdminAuth()
			}
			//isHasAuth:= strings.Contains(b.allowUrl,b.controllerName+"/"+b.actionName)
			noAuth := "ajaxsave/ajaxdel/table/login/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			fmt.Println(b.actionName)
			isNoAuth := strings.Contains(noAuth, b.actionName)
			if isNoAuth == false {
				b.Ctx.WriteString("没有权限")
				b.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}
	if b.userId == 0 && (b.controllerName != "login" && b.actionName != "login") {
		b.redirect(beego.URLFor("LoginController.Login"))
	}

}

//admin权限
func (b *BaseController) AdminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	//b.userName!="admin"{
	//	//非超级管理员
	//	adminAuthIds,_:=m
	//}
}

//登录
func (b *BaseController) LoginAdmin() string {
	admin := models.Admin{}
	errorMsg := ""
	if nil == b.ParseForm(&admin) {
		pwd := admin.Password
		user, err := admin.GetUser(admin.LoginName)
		if err != nil || user.Password != libs.Md5([]byte(pwd+user.Salt)) {
			errorMsg = "帐号或密码错误"
		} else if user.Status == -1 {
			errorMsg = "该帐号已禁用"
		} else {
			user.LastIp = b.getClientIp()
			user.LastLogin = time.Now().Unix()
			user.Update()
			authKey := libs.Md5([]byte(user.LastIp + "|" + user.Password + user.Salt))
			b.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authKey, 7*86400)
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

//ajax返回
func (b *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	b.Data["json"] = out
	b.ServeJSON()
	b.StopRun()
}

// 是否POST提交
func (b *BaseController) isPost() bool {
	return b.Ctx.Request.Method == "POST"
}

//加载模板
func (b *BaseController) disPlay(tpl ...string) {
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = b.controllerName + "/" + b.actionName + ".html"
	}
	b.Layout = "public/layout.html"
	b.TplName = tplName
}
