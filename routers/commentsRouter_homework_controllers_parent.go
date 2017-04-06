package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["homework/controllers/parent:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers/parent:TaskController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers/parent:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers/parent:TaskController"],
		beego.ControllerComments{
			"Show",
			`/:id`,
			[]string{"get"},
			nil})

}
