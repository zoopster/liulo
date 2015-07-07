package routers

import (
	"github.com/astaxie/beego"
	"github.com/dwarvesf/liulo/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",

		// Authentication
		beego.NSNamespace("/users",
			beego.NSRouter("/sign_up", &controllers.UserController{}, "POST:SignUp"),
			beego.NSRouter("/sign_in", &controllers.UserController{}, "POST:SignIn"),
			beego.NSRouter("/", &controllers.UserController{}, "GET:GetCurrentUser"),
		),

		// Events
		beego.NSNamespace("/events",
			beego.NSRouter("/", &controllers.EventController{},
				"POST:Post"),
			beego.NSRouter("/:slug", &controllers.EventController{},
				"GET:GetBySlug"),

			// Sessions
			beego.NSNamespace("/:slug/sessions",
				beego.NSRouter("/", &controllers.SessionController{},
					"GET:GetAll;POST:Post"),
				beego.NSRouter("/:sessionId",
					&controllers.SessionController{},
					"GET:GetById"),

				// Questions
				beego.NSNamespace("/:sessionId/questions",
					beego.NSRouter("/", &controllers.QuestionController{},
						"GET:GetAll;POST:Post"),
					beego.NSRouter("/:questionId", &controllers.QuestionController{},
						"GET:GetById"),
				),
			),
		),
	)

	beego.AddNamespace(ns)
}
