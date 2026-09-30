package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"event-explorer/services"
)

type RedirectController struct {
	beego.Controller
}

func (c *RedirectController) Get() {

	eventID :=
		c.Ctx.Input.Param(":eventId")

	if eventID == "" {

		c.Abort(
			"invalid event id",
		)

		return
	}

	event, err :=
		services.GetEventDetails(
			c.Ctx.Request.Context(),
			eventID,
		)

	if err != nil {

		c.Abort(
			"event not found",
		)

		return
	}

	if !services.IsValidTicketURL(
		event.TicketURL,
	) {

		c.Abort(
			"invalid ticket url",
		)

		return

	}

	c.Redirect(
		event.TicketURL,
		302,
	)

}
