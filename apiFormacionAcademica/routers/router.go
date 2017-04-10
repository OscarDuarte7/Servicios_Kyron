// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Kyron/apiFormacionAcademica/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/nivel_formacion",
			beego.NSInclude(
				&controllers.NivelFormacionController{},
			),
		),

		beego.NSNamespace("/formacion_academica",
			beego.NSInclude(
				&controllers.FormacionAcademicaController{},
			),
		),

		beego.NSNamespace("/institucion",
			beego.NSInclude(
				&controllers.InstitucionController{},
			),
		),

		beego.NSNamespace("/programa",
			beego.NSInclude(
				&controllers.ProgramaController{},
			),
		),

		beego.NSNamespace("/titulo",
			beego.NSInclude(
				&controllers.TituloController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
