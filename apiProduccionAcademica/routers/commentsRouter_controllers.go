package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:CriterioEvaluacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:DatoSubtipoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:EvaluadorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:OpcionDatoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:ProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:SubtipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TipoProduccionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TrProduccionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiProduccionAcademica/controllers:TrProduccionAcademicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

}
