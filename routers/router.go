package routers

import (
	"github.com/astaxie/beego"
	"github.com/dwarvesf/liulo/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		beego.NSNamespace("/events",
			beego.NSInclude(
				&controllers.EventController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
