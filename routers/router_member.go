package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastweb/controllers"
	"github.com/feisuweb/fastweb/filters"
)

func init() {
	//会员服务
	beego.InsertFilter("/member", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/", &controllers.MemberHandle{}, "*:Index")

	beego.InsertFilter("/member/activate", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/activate", &controllers.MemberHandle{}, "*:Activate")

	beego.Router("/member/buy", &controllers.MemberHandle{}, "*:Buy")

	beego.InsertFilter("/member/upgrade", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/upgrade", &controllers.MemberHandle{}, "*:Upgrade")

	beego.InsertFilter("/member/createvip", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/createvip", &controllers.MemberHandle{}, "*:CreateVip")

	beego.InsertFilter("/member/upgradevip", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/upgradevip", &controllers.MemberHandle{}, "*:UpgradeVip")

}
