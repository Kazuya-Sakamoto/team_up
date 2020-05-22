package routers

import (
	"app/server_side/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/signup",
			beego.NSInclude(
				&controllers.SignupController{},
			),
		),
		beego.NSNamespace("/job",
			beego.NSInclude(
				&controllers.JobController{},
			),
		),
		// beego.NSNamespace("/access_right",
		// 	beego.NSInclude(
		// 		&controllers.AccessRightController{},
		// 	),
		// ),
		// beego.NSNamespace("/role",
		// 	beego.NSInclude(
		// 		&controllers.RoleController{},
		// 	),
		// ),
	)
	beego.AddNamespace(ns)
}
