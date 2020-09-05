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
		beego.NSNamespace("/programing_language",
			beego.NSInclude(
				&controllers.ProgramingLanguageController{},
			),
		),
		beego.NSNamespace("/programing_framework",
			beego.NSInclude(
				&controllers.ProgramingFrameworkController{},
			),
		),
		beego.NSNamespace("/skill",
			beego.NSInclude(
				&controllers.SkillController{},
			),
		),
		// beego.NSNamespace("/communication_tool",
		// 	beego.NSInclude(
		// 		&controllers.SkillController{},
		// 	),
		// ),
		// beego.NSNamespace("/position_tag",
		// 	beego.NSInclude(
		// 		&controllers.PositionTagController{},
		// 	),
		// ),
		beego.NSNamespace("/job_status",
			beego.NSInclude(
				&controllers.JobStatusController{},
			),
		),
		beego.NSNamespace("/individual_portfolio",
			beego.NSInclude(
				&controllers.IndividualPortfolioController{},
			),
		),
		beego.NSNamespace("/chat_message",
			beego.NSInclude(
				&controllers.ChatMessageController{},
			),
		),
		beego.NSNamespace("/favorite_job",
			beego.NSInclude(
				&controllers.FavoriteJobController{},
			),
		),
		beego.NSNamespace("/apply_job",
			beego.NSInclude(
				&controllers.ApplyJobController{},
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
