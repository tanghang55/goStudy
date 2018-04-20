package routers

import (
	"goStudy/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{},"*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
}
