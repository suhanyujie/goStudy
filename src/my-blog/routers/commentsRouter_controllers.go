package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["my-blog/controllers:TopicController"] = append(beego.GlobalControllerRouter["my-blog/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/v1/article/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-blog/controllers:TopicController"] = append(beego.GlobalControllerRouter["my-blog/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/v1/article/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
