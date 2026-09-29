package main

import (
	_ "event-explorer/routers"

	"event-explorer/config"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	config.Load()

	beego.Run()

}
