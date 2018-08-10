package controllers

import (
	"github.com/astaxie/beego"
	"my-blog/models"
)

type CateController struct {
	beego.Controller
}

func (this *CateController) Get() {
	op := this.Input().Get("op")
	name := this.Input().Get("name")
	switch op {
	case "add":
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
	case "del":
		err := models.DeleteCategory("2")
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
	}
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		return
	}
}
