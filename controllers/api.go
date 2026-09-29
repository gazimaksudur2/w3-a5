package controllers

import (
	"event-explorer/services"

	beego "github.com/beego/beego/v2/server/web"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) Autocomplete(){

	input:=c.GetString("input")

	session:=c.GetString("sessionToken")


	results,err := services.Autocomplete(
		input,
		session,
	)


	if err!=nil{

		c.Data["json"]=map[string]string{
			"error":err.Error(),
		}

		c.ServeJSON()
		return

	}



	c.Data["json"]=results

	c.ServeJSON()

}

func (c *APIController) GetLocation() {

	c.Data["json"] = map[string]string{

		"message": "location endpoint working",
	}

	c.ServeJSON()

}
