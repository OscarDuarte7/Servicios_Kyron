// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Kyron/apiProduccionAcademica/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_produccion",
			beego.NSInclude(
				&controllers.TipoProduccionController{},
			),
		),

		beego.NSNamespace("/evaluador",
			beego.NSInclude(
				&controllers.EvaluadorController{},
			),
		),

		beego.NSNamespace("/dato",
			beego.NSInclude(
				&controllers.DatoController{},
			),
		),

		beego.NSNamespace("/criterio_evaluacion",
			beego.NSInclude(
				&controllers.CriterioEvaluacionController{},
			),
		),

		beego.NSNamespace("/produccion_academica",
			beego.NSInclude(
				&controllers.ProduccionAcademicaController{},
			),
		),

		beego.NSNamespace("/dato_subtipo",
			beego.NSInclude(
				&controllers.DatoSubtipoController{},
			),
		),

		beego.NSNamespace("/dato_produccion",
			beego.NSInclude(
				&controllers.DatoProduccionController{},
			),
		),

		beego.NSNamespace("/subtipo_produccion",
			beego.NSInclude(
				&controllers.SubtipoProduccionController{},
			),
		),

		beego.NSNamespace("/opcion_dato",
			beego.NSInclude(
				&controllers.OpcionDatoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
