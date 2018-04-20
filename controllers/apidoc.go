package controllers

type ApiDocController struct {
	BaseController
}

func (a *ApiDocController) Index() {
	a.TplName = "apidoc/index.html"

}
