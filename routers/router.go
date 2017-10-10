package routers

import (
	"accept_token/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
}
