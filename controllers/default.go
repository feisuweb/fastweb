package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "首页"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.TplName = "index.html"
}

func (c *MainController) Login() {
	captchaId := captcha.NewLen(6) //验证码长度为6

	// id, value := this.GetString("captcha_id"), this.GetString("captcha")
	// b := captcha.VerifyString(id, value) //验证码校验
	// this.Ctx.WriteString(strconv.FormatBool(b))

	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "登录"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.Data["CaptchaId"] = captchaId
	c.TplName = "login.html"
}

func (c *MainController) LoginCheck() {
	captchaId := captcha.NewLen(6) //验证码长度为6

	// id, value := this.GetString("captcha_id"), this.GetString("captcha")
	// b := captcha.VerifyString(id, value) //验证码校验
	// this.Ctx.WriteString(strconv.FormatBool(b))

	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "登录"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.Data["CaptchaId"] = captchaId
	c.TplName = "login.html"
}

func (c *MainController) SignUp() {
	captchaId := captcha.NewLen(6) //验证码长度为6
	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "注册"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.Data["CaptchaId"] = captchaId
	c.TplName = "signup.html"
}

func (c *MainController) SignUpCheck() {
	captchaId := captcha.NewLen(6) //验证码长度为6
	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "注册"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.Data["CaptchaId"] = captchaId
	c.TplName = "signup.html"
}

func (c *MainController) Explore() {
	c.Data["SiteName"] = "微信采集服务"
	c.Data["Title"] = "浏览"
	c.Data["Keywords"] = "astaxie@gmail.com"
	c.Data["Description"] = "astaxie@gmail.com"
	c.Data["Author"] = "astaxie@gmail.com"
	c.Data["XsrfToken"] = c.XSRFToken()
	c.TplName = "explore.html"
}
