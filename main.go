package main

import (
	_ "firmware/controllers"
	_ "firmware/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()

}
