package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"event-explorer/controllers"
)

func init() {

	beego.Router(
		"/",
		&controllers.HomeController{},
	)

	beego.Router(
		"/events",
		&controllers.EventController{},
	)

	beego.Router(
		"/events/:eventId",
		&controllers.EventController{},
	)

	beego.Router(
		"/api/locations/autocomplete",
		&controllers.APIController{},
		"get:Autocomplete",
	)

	beego.Router(
		"/api/locations/:placeId",
		&controllers.APIController{},
		"get:GetLocation",
	)

}
