package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastweb/controllers"
	"github.com/feisuweb/fastweb/filters"
)

func init() {
	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //注册验证码服务，验证码图片的宽高为240 x 80
	beego.Router("/", &controllers.MainController{})
	beego.Router("/explore", &controllers.MainController{}, "*:Explore")

	beego.Router("/login", &controllers.MemberHandle{}, "get:GetLogin")
	beego.Router("/login", &controllers.MemberHandle{}, "post:PostLogin")

	beego.Router("/register", &controllers.MemberHandle{}, "get:GetRegister")
	beego.Router("/register", &controllers.MemberHandle{}, "post:PostRegister")
	//会员服务

	beego.Router("/member/buy", &controllers.MemberHandle{}, "*:Buy")

	beego.InsertFilter("/member/upgrade", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/upgrade", &controllers.MemberHandle{}, "*:Upgrade")

	beego.InsertFilter("/member/", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/", &controllers.MemberHandle{}, "*:Index")

	beego.InsertFilter("/member/createvip", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/createvip", &controllers.MemberHandle{}, "*:CreateVip")

	beego.InsertFilter("/member/upgradevip", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/upgradevip", &controllers.MemberHandle{}, "*:UpgradeVip")

	beego.InsertFilter("/member/profile", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/profile", &controllers.MemberHandle{}, "*:Profile")

	beego.Router("/member/logout", &controllers.MemberHandle{}, "*:Logout")

	beego.InsertFilter("/member/changepassword", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/changepassword", &controllers.MemberHandle{}, "get:GetChangePassword")

	beego.InsertFilter("/member/changepassword", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/changepassword", &controllers.MemberHandle{}, "post:PostChangePassword")

}
