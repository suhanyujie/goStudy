package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Get() {

	beego.SetLevel(beego.LevelDebug)
	// beego.Alert("this is alert")
	c.Data["Email"] = "suhanyujie@qq.com"
	c.Data["Html1"] = "<div>Hello beego</div>"
	c.Data["IsIndex"] = true
	c.Data["HasLogin"] = CheckAcount(c.Ctx)
	c.TplName = "Index/home.html"
}
