package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	this.TplName = "article/index.html"
}
