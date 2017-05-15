package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastweb/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //注册验证码服务，验证码图片的宽高为240 x 80
	beego.Router("/", &controllers.MainController{})

	beego.Router("/explore", &controllers.MainController{}, "*:Explore")

	beego.Router("/login", &controllers.MemberHandle{}, "get:GetLogin")
	beego.Router("/login", &controllers.MemberHandle{}, "post:PostLogin")
	beego.Router("/register", &controllers.MemberHandle{}, "get:GetRegister")
	beego.Router("/register", &controllers.MemberHandle{}, "post:PostRegister")
	beego.Router("/logout", &controllers.MemberHandle{}, "*:Logout")
}
