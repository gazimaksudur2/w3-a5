package controllers


import (
	beego "github.com/beego/beego/v2/server/web"
)



type EventController struct{

	beego.Controller

}



func (c *EventController)Get(){


c.TplName="listing.tpl"


}



func (c *EventController)GetDetails(){


c.TplName="details.tpl"


}