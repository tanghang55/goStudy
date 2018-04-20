package controllers

type HomeController struct{
	BaseController
}

func (h *HomeController) Index() {
	h.Data["pageTitle"] = "系统首页"
	h.TplName = "public/main.html"
}

func (h *HomeController) Start() {
	h.Data["pageTitle"] = "控制面板"
	h.disPlay()
}