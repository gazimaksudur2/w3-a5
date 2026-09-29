package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) Autocomplete() {

	c.Data["json"] = map[string]string{

		"message": "autocomplete endpoint working",
	}

	c.ServeJSON()

}

func (c *APIController) GetLocation() {

	c.Data["json"] = map[string]string{

		"message": "location endpoint working",
	}

	c.ServeJSON()

}
