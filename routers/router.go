package routers


import (
	beego "github.com/beego/beego/v2/server/web"

	"event-explorer/controllers"
)


func init(){


beego.Router(
"/",
&controllers.HomeController{},
)


}