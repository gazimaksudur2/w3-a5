package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"event-explorer/services"
)

type EventController struct {
	beego.Controller
}

func (c *EventController) Get() {

	city := c.GetString("city")

	country := c.GetString("countryCode")

	if city == "" || country == "" {
		c.Data["Error"] = "Please select a city"
		c.TplName = "listing.tpl"
		return

	}

	provider := services.TicketmasterProvider{}
	events, err := services.GetCityEvents(
		c.Ctx.Request.Context(),
		provider,
		city,
		country,
	)

	if err != nil {
		c.Data["Error"] = err.Error()

	}

	c.Data["City"] = city

	c.Data["Music"] = events.Music

	c.Data["Sports"] = events.Sports

	c.TplName = "listing.tpl"

}
