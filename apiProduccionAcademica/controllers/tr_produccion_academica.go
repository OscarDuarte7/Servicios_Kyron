package controllers

import (
	"Kyron/apiProduccionAcademica/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// oprations for TrProduccionAcademicaController
type TrProduccionAcademicaController struct {
	beego.Controller
}

func (c *TrProduccionAcademicaController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// AddTransaccionProduccionAcademica ...

func (c *TrProduccionAcademicaController) Post() {

	var v models.TrProduccionAcademica
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if _, err = models.AddTransaccionProduccionAcademica(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
