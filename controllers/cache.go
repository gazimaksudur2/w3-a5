package controllers


import (
	"event-explorer/services"

	beego "github.com/beego/beego/v2/server/web"
)



type CacheController struct {
	beego.Controller
}



func (c *CacheController) Invalidate(){


	city :=
		c.GetString("city")


	country :=
		c.GetString("country")


	category :=
		c.GetString("category")



	services.EventCacheInstance.Invalidate(
		city,
		country,
		category,
	)



	c.Data["json"] = map[string]string{
		"message":"cache invalidated",
	}


	c.ServeJSON()

}