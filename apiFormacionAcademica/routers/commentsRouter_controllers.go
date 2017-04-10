package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"] = append(beego.GlobalControllerRouter["Kyron/apiFormacionAcademica/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
