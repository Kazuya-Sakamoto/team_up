package routers

import (
	"app/server_side/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api/vi") // beego.Namespace("/user",
	// 	beego.NSInclude(
	// 		&controllers.UserController{},
	// 	),
	// ),

	beego.AddNamespace(ns)
}
