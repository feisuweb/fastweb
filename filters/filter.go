package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/feisuweb/fastweb/models"
)

func IsLogin(ctx *context.Context) (bool, models.Member) {
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	var memberInfo models.Member
	if flag {
		flag, memberInfo = models.FindMemberByToken(token)
	}
	return flag, memberInfo
}

var FilterUser = func(ctx *context.Context) {
	ok, LoginMember := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	}
	//用户激活判断
	if LoginMember.MemberActivated == 0 && ctx.Request.RequestURI != "/activate" && ctx.Request.RequestURI != "/active" {
		ctx.Redirect(302, "/member/activate")
	}
}
