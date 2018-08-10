package main

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"my-blog/controllers"
	"my-blog/models"
	_ "my-blog/routers"
)

func init() {
	models.MyRegister()
}

func main() {
	orm.Debug = true

	// o := orm.NewOrm()
	// var maps []orm.Params
	// _, err := o.Raw("SELECT * FROM tests").Values(&maps)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, term := range maps {
	// 	fmt.Println(term["content"])
	// }

	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})

	beego.Run()
}
