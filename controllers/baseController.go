package controllers

import (
	"github.com/astaxie/beego"
	"goStudy/models"
)

type BaseController struct{
	beego.Controller
	admin  models.Admin
}
