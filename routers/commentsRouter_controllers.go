package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["accept_token/controllers:MainController"] = append(beego.GlobalControllerRouter["accept_token/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["accept_token/controllers:MainController"] = append(beego.GlobalControllerRouter["accept_token/controllers:MainController"],
		beego.ControllerComments{
			Method: "Access",
			Router: `/access`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
