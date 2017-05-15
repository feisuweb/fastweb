package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	//"github.com/dchest/captcha"
	"github.com/feisuweb/fastweb/libs/notify"
	"github.com/feisuweb/fastweb/models"
	"strings"
	"time"
)

///前台页面handle
type SettingHandle struct {
	baseController
}

//会员激活账号
func (this *SettingHandle) Activate() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_activate_resend.html"
}

//会员个人资料
func (this *SettingHandle) GetProfile() {

	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "注册会员"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_profile.html"
}

func (this *SettingHandle) ShowRegisterError(errorMsg string) {
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_register.html"
}

//会员登录
func (this *SettingHandle) PostRegister() {
	var (
		minfo    *models.Member = new(models.Member)
		err      bool
		mobile   string
		email    string
		password string
		ip       string
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "会员登录"
	if !models.ValidateMobile(mobile) {
		this.ShowRegisterError("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.ShowRegisterError("请填写正确格式的邮箱！")
		return
	}
	if len(password) == 0 {
		this.ShowRegisterError("请填写密码！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	err = minfo.FindMemberByMobileOrEmail(mobile, email)
	if err {
		//如果查询到用户已经存在，则提示用户已经存在了。
		this.ShowRegisterError("手机号或者邮箱已经注册过会员账号。")
		return
	}

	//注册账号信息
	minfo.Email = email
	minfo.Password = password
	minfo.Mobile = mobile
	minfo.Nickname = mobile
	minfo.MemberName = mobile
	minfo.RegisterIp = ip
	minfo.IsVip = 0
	minfo.IsValidateMobile = 0
	minfo.IsValidateEmail = 0
	minfo.Points = 0
	minfo.Money = 0
	ret := minfo.Register()

	if ret {
		//注册成功，跳转到会员首页
		this.Redirect("/member", 302)
		return
	} else {
		this.ShowRegisterError("账号注册失败！")
		return
	}
}

//会员头像
func (this *SettingHandle) GetAvatar() {

	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "注册会员"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_avatar.html"
}

//修改密码
func (this *SettingHandle) GetSettingPassword() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_password.html"
}

//修改密码-POST
func (this *SettingHandle) PostSettingPassword() {
	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		minfo     *models.Member        = new(models.Member)
		msloginfo *models.MemberSafeLog = new(models.MemberSafeLog)
		err       bool
		ip        string
	)
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	oldPassword := strings.TrimSpace(this.GetString("oldpassword"))
	newPassword := strings.TrimSpace(this.GetString("newpassword"))
	err = minfo.ChangePassword(LoginMember.Id, oldPassword, newPassword)

	if err {
		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改成功！")
		//发送密码修改通知给会员
		t := time.Now().Format("2006-01-02 15:04:05")
		var ni notify.NotifyInfo
		ni.MemberName = LoginMember.MemberName
		ni.MemberEmail = LoginMember.Email
		ni.MemberMobile = LoginMember.Mobile
		ni.MemberWeixinOpenId = LoginMember.WeixinOpenId
		ni.ChangePasswordTime = t
		ni.ChangePasswordIp = ip
		ni.ChangePasswordNewPassword = newPassword

		go notify.SendToMemberPasswordChangedNotify(&ni)

		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			"", -1,
			"/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		this.Redirect("/login", 302)
		return
	} else {

		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改失败！")
	}

	this.Data["ErrorMsg"] = "修改密码失败！"
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_password.html"
}

//修改email
func (this *SettingHandle) GetSettingEmail() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_email.html"
}

//修改email-POST
func (this *SettingHandle) PostSettingEmail() {
	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		minfo     *models.Member        = new(models.Member)
		msloginfo *models.MemberSafeLog = new(models.MemberSafeLog)
		err       bool
		ip        string
	)
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	oldPassword := strings.TrimSpace(this.GetString("oldpassword"))
	newPassword := strings.TrimSpace(this.GetString("newpassword"))
	err = minfo.ChangePassword(LoginMember.Id, oldPassword, newPassword)

	if err {
		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改成功！")
		//发送密码修改通知给会员
		t := time.Now().Format("2006-01-02 15:04:05")
		var ni notify.NotifyInfo
		ni.MemberName = LoginMember.MemberName
		ni.MemberEmail = LoginMember.Email
		ni.MemberMobile = LoginMember.Mobile
		ni.MemberWeixinOpenId = LoginMember.WeixinOpenId
		ni.ChangePasswordTime = t
		ni.ChangePasswordIp = ip
		ni.ChangePasswordNewPassword = newPassword

		go notify.SendToMemberPasswordChangedNotify(&ni)

		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			"", -1,
			"/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		this.Redirect("/login", 302)
		return
	} else {

		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改失败！")
	}

	this.Data["ErrorMsg"] = "修改密码失败！"
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_email.html"
}

//修改安全设置
func (this *SettingHandle) GetSettingSecurity() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_security.html"
}

//修改安全设置-POST
func (this *SettingHandle) PostSettingSecurity() {
	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		minfo     *models.Member        = new(models.Member)
		msloginfo *models.MemberSafeLog = new(models.MemberSafeLog)
		err       bool
		ip        string
	)
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	oldPassword := strings.TrimSpace(this.GetString("oldpassword"))
	newPassword := strings.TrimSpace(this.GetString("newpassword"))
	err = minfo.ChangePassword(LoginMember.Id, oldPassword, newPassword)

	if err {
		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改成功！")
		//发送密码修改通知给会员
		t := time.Now().Format("2006-01-02 15:04:05")
		var ni notify.NotifyInfo
		ni.MemberName = LoginMember.MemberName
		ni.MemberEmail = LoginMember.Email
		ni.MemberMobile = LoginMember.Mobile
		ni.MemberWeixinOpenId = LoginMember.WeixinOpenId
		ni.ChangePasswordTime = t
		ni.ChangePasswordIp = ip
		ni.ChangePasswordNewPassword = newPassword

		go notify.SendToMemberPasswordChangedNotify(&ni)

		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			"", -1,
			"/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		this.Redirect("/login", 302)
		return
	} else {

		msloginfo.AddSafeLog(LoginMember, "changepassword", ip, "密码修改失败！")
	}

	this.Data["ErrorMsg"] = "修改密码失败！"
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_member_settings_layout.html"
	this.TplName = "member/settings/_security.html"
}
