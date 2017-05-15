package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastweb/controllers"
	"github.com/feisuweb/fastweb/filters"
)

func init() {

	//设置模块
	beego.InsertFilter("/member/settings", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings", &controllers.SettingHandle{}, "*:GetProfile")

	beego.InsertFilter("/member/settings/profile", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/profile", &controllers.SettingHandle{}, "*:GetProfile")

	beego.InsertFilter("/member/settings/avatar", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/avatar", &controllers.SettingHandle{}, "*:GetAvatar")

	beego.InsertFilter("/member/settings/password", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/password", &controllers.SettingHandle{}, "get:GetSettingPassword")

	beego.InsertFilter("/member/settings/password", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/password", &controllers.SettingHandle{}, "post:PostSettingPassword")

	beego.InsertFilter("/member/settings/email", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/email", &controllers.SettingHandle{}, "get:GetSettingEmail")

	beego.InsertFilter("/member/settings/email", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/email", &controllers.SettingHandle{}, "post:PostSettingEmail")

	beego.InsertFilter("/member/settings/security", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/security", &controllers.SettingHandle{}, "get:GetSettingSecurity")

	beego.InsertFilter("/member/settings/security", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/member/settings/security", &controllers.SettingHandle{}, "post:PostSettingSecurity")
}
