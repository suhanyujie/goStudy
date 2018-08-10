package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["test1"] = beego.AppConfig.String("email")
	this.Data["HasLogin"] = CheckAcount(this.Ctx)

	this.TplName = "Login/login.html"
}

func (this *LoginController) Post() {
	email := this.Input().Get("email")
	passwd := this.Input().Get("passwd")
	isRemember := this.Input().Get("remember") == "on"
	if email == beego.AppConfig.String("email") &&
		passwd == beego.AppConfig.String("passwd") {
		maxAge := 0
		if isRemember {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("email", email, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func CheckAcount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("email")
	fmt.Println(ck)
	if err != nil {
		return false
	}
	return true
}
