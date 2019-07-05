// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ApiGuide/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/get_user_detail", &controllers.UserController{}, "get:GetUserInfo"),
		),
		beego.NSNamespace("/project",
			beego.NSRouter("/get_all_project", &controllers.ProjectController{}, "get:GetAllPjc"),
			beego.NSRouter("/add_project", &controllers.ProjectController{}, "post:AdPro"),
			beego.NSRouter("/delete", &controllers.ProjectController{}, "delete:DelPjt"),
		),
		beego.NSNamespace("/modular",
			beego.NSRouter("/get_all_modular", &controllers.ModularController{}, "get:GetAllModular"),
			beego.NSRouter("/add_modular", &controllers.ModularController{}, "post:AddModular"),
		),
		beego.NSNamespace("/interface",
			beego.NSRouter("/add_interface", &controllers.InterfaceController{}, "post:AddInterface"),
			beego.NSRouter("/get_interface", &controllers.InterfaceController{}, "get:GetInterface"),
		),
	)
	beego.AddNamespace(ns)
}
