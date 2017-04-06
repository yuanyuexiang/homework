package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"PutComment",
			`/:id/comment`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"PutFlowers",
			`/:id/flowers`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"PutPraises",
			`/:id/praises`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:SolutionController"] = append(beego.GlobalControllerRouter["homework/controllers:SolutionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:StatisticController"] = append(beego.GlobalControllerRouter["homework/controllers:StatisticController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:StatisticController"] = append(beego.GlobalControllerRouter["homework/controllers:StatisticController"],
		beego.ControllerComments{
			"PutStatisticComment",
			`/:id/comment`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers:TaskController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers:TaskController"],
		beego.ControllerComments{
			"Show",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers:TaskController"],
		beego.ControllerComments{
			"Create",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["homework/controllers:TaskController"] = append(beego.GlobalControllerRouter["homework/controllers:TaskController"],
		beego.ControllerComments{
			"Destory",
			`/:id`,
			[]string{"delete"},
			nil})

}
