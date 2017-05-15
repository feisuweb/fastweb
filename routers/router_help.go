package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastweb/controllers"
)

func init() {
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
}
