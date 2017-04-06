// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"homework/controllers"
	"homework/controllers/parent"

	"github.com/astaxie/beego"
)

func init() {
	ns_t := beego.NewNamespace("homework/v1/teacher",
		beego.NSNamespace("/tasks",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
		beego.NSNamespace("/solutions",
			beego.NSInclude(
				&controllers.SolutionController{},
			),
		),
		beego.NSNamespace("/statistics",
			beego.NSInclude(
				&controllers.StatisticController{},
			),
		),
	)
	ns_p := beego.NewNamespace("homework/v1/parent",
		beego.NSNamespace("/tasks",
			beego.NSInclude(
				&parent.TaskController{},
			),
		),
		beego.NSNamespace("/solutions",
			beego.NSInclude(
				&controllers.SolutionController{},
			),
		),
		beego.NSNamespace("/statistics",
			beego.NSInclude(
				&controllers.StatisticController{},
			),
		),
	)
	beego.AddNamespace(ns_t)
	beego.AddNamespace(ns_p)
}
