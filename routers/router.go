package routers

import (
	"firmware/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/firmware/version", &controllers.FirmController{}, "post:FindNewVer") //查询固件
	beego.Router("/firmware/publish", &controllers.FirmController{}, "post:NewVer")     //发布固件
	beego.Router("firmware/download", &controllers.FirmController{}, "post:DownloadFirm")
}
