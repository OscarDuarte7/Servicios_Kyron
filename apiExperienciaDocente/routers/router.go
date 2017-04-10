// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Kyron/apiExperienciaDocente/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/experiencia_docente",
			beego.NSInclude(
				&controllers.ExperienciaDocenteController{},
			),
		),

		beego.NSNamespace("/cursos",
			beego.NSInclude(
				&controllers.CursosController{},
			),
		),

		beego.NSNamespace("/tipo_dedicacion",
			beego.NSInclude(
				&controllers.TipoDedicacionController{},
			),
		),

		beego.NSNamespace("/institucion",
			beego.NSInclude(
				&controllers.InstitucionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
