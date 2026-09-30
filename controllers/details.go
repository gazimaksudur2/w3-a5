package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"event-explorer/services"
)

type DetailsController struct {
	beego.Controller
}

func (c *DetailsController) Get() {

	eventID := c.Ctx.Input.Param(":eventId")

	if eventID == "" {

		c.Data["Error"] = "Invalid event ID"
		c.TplName = "details.tpl"
		return

	}

	event, err := services.GetEventDetails(c.Ctx.Request.Context(), eventID)

	if err != nil {

		c.Data["Error"] = err.Error()
		c.TplName = "details.tpl"
		return

	}

	c.Data["Event"] = event

	c.TplName = "details.tpl"

}
