package routers

import (
	"github.com/astaxie/beego"
	"my-blog/controllers"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Get("/test/a/test1", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob...."))
	})

	beego.Include(&controllers.TopicController{})
}

