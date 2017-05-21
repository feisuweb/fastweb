package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastweb/models"
	"strings"
	//"time"
)

///前台页面handle
type MemberHandle struct {
	baseController
}

//会员首页
func (this *MemberHandle) Index() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "会员首页"
	this.Layout = "layout/_member_layout.html"
	this.TplName = "member/_index.html"
}

//会员激活账号
func (this *MemberHandle) Activate() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Layout = "layout/_member_layout.html"
	this.TplName = "member/_activate_resend.html"
}

//会员注册
func (this *MemberHandle) GetRegister() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "注册会员"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "member/_register.html"
}

func (this *MemberHandle) ShowRegisterError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "member/_register.html"
}

//会员登录
func (this *MemberHandle) PostRegister() {
	var (
		minfo    *models.Member = new(models.Member)
		err      bool
		mobile   string
		email    string
		password string
		ip       string
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	id, value := this.GetString("captcha_id"), this.GetString("captcha")
	b := captcha.VerifyString(id, value) //验证码校验
	if !b {
		this.ShowRegisterError("验证码错误！")
		return
	}
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
func (this *MemberHandle) ShowLoginError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Data["PageTitle"] = "会员登录 - 登录错误"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "member/_login.html"
}

//会员登录
func (this *MemberHandle) GetLogin() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "会员登录"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "member/_login.html"
}

//会员登录
func (this *MemberHandle) PostLogin() {

	var (
		minfo      *models.Member = new(models.Member)
		err        bool
		memberName string
		password   string
		ip         string
	)
	memberName = strings.TrimSpace(this.GetString("account"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "会员登录"
	if len(memberName) == 0 {
		this.ShowLoginError("请填写用户名！")
		return
	}
	if len(password) == 0 {
		this.ShowLoginError("请填写密码！")
		return
	}

	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	err = minfo.Login(memberName, password, ip)
	if err {
		//登录成功
		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			minfo.Token, 30*24*60*60, "/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		mid2 := fmt.Sprintf("%d", minfo.Id)
		this.Ctx.SetCookie("member_id", mid2)
		this.Redirect("/member/", 302)
		return
	} else {
		//登录失败
		this.ShowLoginError("账号或者密码错误")
		return
	}
}

//会员退出登录
func (this *MemberHandle) Logout() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.SetSecureCookie(
		beego.AppConfig.String("cookie.secure"),
		beego.AppConfig.String("cookie.token"),
		"", -1,
		"/",
		beego.AppConfig.String("cookie.domain"),
		false,
		true)
	this.Redirect("/login", 302)
}

//会员找回密码
func (this *MemberHandle) FindPassword() {

	this.Ctx.Output.Header("Cache-Control", "public")

	this.Layout = "layout/_member_layout.html"

	this.TplName = "member/_findpassword.html"
}

//购买会员服务
func (this *MemberHandle) Buy() {

	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		memberInfo  models.Member
		memberOrder models.MemberOrder
	)

	memberList := memberOrder.GetLastList(6)

	this.Data["PageTitle"] = "购买会员"
	this.Data["memberlist"] = memberList

	this.Data["memberInfo"] = memberInfo

	this.Layout = "layout/_member_layout.html"

	this.TplName = "member/_buy.html"
}

//升级会员服务
func (this *MemberHandle) Upgrade() {

	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		memberInfo  models.Member
		memberOrder models.MemberOrder
	)

	memberList := memberOrder.GetLastList(6)

	this.Data["PageTitle"] = "升级会员"
	this.Data["memberlist"] = memberList

	this.Data["memberInfo"] = memberInfo

	this.Layout = "layout/_member_layout.html"

	this.TplName = "member/_upgrade.html"
}

//购买VIP会员
func (this *MemberHandle) CreateVip() {

	var (
		info        *models.MemberOrder = new(models.MemberOrder)
		productInfo *models.MemberType  = new(models.MemberType)
		minfo       *models.Member      = new(models.Member)
		payinfo     *models.PayLog      = new(models.PayLog)
		//agentInfo      *models.Agent       = new(models.Agent)
		err            bool
		orderNo        string
		url            string
		member_id      int64
		member_type_id int64
		mobile         string
		email          string
		password       string
		ip             string
	)

	member_id, _ = this.GetInt64("member_id")
	member_type_id, _ = this.GetInt64("vip_type")
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))

	if !models.ValidateMobile(mobile) {
		this.Abort("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.Abort("请填写正确格式的邮箱！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	err = minfo.FindMemberByMobileAndEmail(mobile, email)
	if err {
		//如果查询到用户已经存在，则
		member_id = minfo.Id
	} else {
		//注册账号信息
		//默认以邮箱和手机号注册一个用户，用户密码是随机数
		//memberName string, password string, mobile string, email string, ip string
		//ipResult := models.TabaoAPI(ip)
		minfo.Email = email
		minfo.Password = password
		minfo.Mobile = mobile
		minfo.Nickname = "会员" + mobile
		minfo.MemberName = mobile
		minfo.RegisterIp = ip
		minfo.IsVip = 0
		minfo.IsValidateMobile = 0
		minfo.IsValidateEmail = 0
		minfo.Points = 0
		minfo.Money = 0
		//如果有代理商信息
		//if agent_id > 0 && len(agent_mobile) > 0 {
		//minfo.Id = ""
		//minfo.AgentMobile = ""

		//}
		err := minfo.Register()
		if err {
			member_id = minfo.Id
			//更新代理商数据
			//如果有代理商信息
			//if agent_id > 0 && len(agent_mobile) > 0 {
			//agentInfo.Id = ""
			//agentInfo.AddMemberCount()

			//}
		}
	}
	//根据产品ID查询产品信息

	err = productInfo.FindMemberTypeById(member_type_id)

	if !err {
		this.Abort("会员类型信息有误，请查验后再提交")
	}

	//如果是VIP会员，则直接判断
	// if minfo.CheckVip(minfo.Id) {
	// 	this.Abort("已经是VIP，无需再次购买！")
	// }
	//判断之前是否已经购买过，购买过则无需再次购买
	//订单创建流程开始
	//获取随机订单号
	orderNo = info.GetRandOrderNo()
	//订单创建
	info.OrderNo = orderNo
	info.ProductId = member_type_id
	info.ProductName = productInfo.Name
	info.MemberId = member_id

	info.FromPlatform = "pc"
	info.FromChannel = "direct"
	info.FromChannelTag = "codeshop.com"

	// info.RecommendCode = agentInfo.RecommendCode
	// info.AgentId = agentInfo.Id
	// info.AgentName = agentInfo.AgentName
	// info.AgentWeixinOpenId = agentInfo.WeixinOpenId
	// info.AgentWeixin = agentInfo.Weixin
	// info.AgentEmail = agentInfo.Email
	// info.AgentMobile = agentInfo.Mobile

	info.MemberName = minfo.Nickname
	info.MemberMobile = minfo.Mobile
	info.MemberEmail = minfo.Email
	info.MemberWeixin = minfo.Weixin
	info.MemberWeixinOpenId = minfo.WeixinOpenId

	info.CommissionAmount = 0
	info.Count = 1
	if minfo.CheckVip(member_id) {
		//VIP 会员，采用会员价购买
		info.Price = productInfo.Price
		info.Discount = 0
	} else {
		//普通会员，采用普通价格购买
		info.Price = productInfo.Price
		info.Discount = 0
	}
	info.PayAmount = info.Price
	info.Amount = info.Price
	info.IsSend = 0
	info.Status = 0
	//创建订单
	orderId, oerr := info.CreateOrder()
	if oerr == false {
		this.Abort("会员订单创建失败")
	}

	//创建微信支付记录
	payinfo.OrderId = orderId
	payinfo.OrderNo = info.OrderNo
	payinfo.PayType = 1 //消费
	payinfo.OrderType = "member"
	payinfo.MemberId = member_id
	payinfo.AgentId = info.AgentId
	payinfo.MemberName = info.MemberName
	payinfo.MemberMobile = mobile
	payinfo.MemberEmail = email
	payinfo.MemberWeixin = info.MemberWeixin
	payinfo.Amount = info.Amount
	payinfo.Discount = info.Discount
	payinfo.PayAmount = info.PayAmount
	payinfo.PayMethod = "weixin"
	payinfo.PayBody = "购买会员服务" + info.ProductName + "-优品源码网"
	payinfo.ProductId = info.ProductId
	payinfo.PayStatus = 0
	payinfo.Status = 0
	payinfo.Insert()

	//增加代理商VIP会员数
	//if agent_id > 0 && len(agent_mobile) > 0 {
	//	agentInfo.Id = ""
	//	agentInfo.AddVipMemberCount()

	//}

	url = site_pay_scan_url + "?orderno=" + orderNo
	if info.PayAmount > 0 {
		url = site_pay_scan_url + "?orderno=" + orderNo
	} else {
		//直接跳转会员详细页面
		url = fmt.Sprintf("/member/profile/%d")
	}
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	mid3 := fmt.Sprintf("%d", member_id)
	this.Ctx.SetCookie("member_id", mid3)

	this.SetSecureCookie(
		beego.AppConfig.String("cookie.secure"),
		beego.AppConfig.String("cookie.token"),
		minfo.Token, 30*24*60*60, "/",
		beego.AppConfig.String("cookie.domain"),
		false,
		true)

	this.Redirect(url, 302)
	return
}

//升级VIP会员
func (this *MemberHandle) UpgradeVip() {

	var (
		info        *models.MemberOrder = new(models.MemberOrder)
		productInfo *models.MemberType  = new(models.MemberType)
		minfo       *models.Member      = new(models.Member)
		payinfo     *models.PayLog      = new(models.PayLog)
		//agentInfo      *models.Agent       = new(models.Agent)
		err            bool
		orderNo        string
		url            string
		member_id      int64
		member_type_id int64
		mobile         string
		email          string
		//ip             string
	)
	member_id = LoginMember.Id
	member_type_id, _ = this.GetInt64("vip_type")
	mobile = LoginMember.Mobile
	email = LoginMember.Email

	//ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	minfo.FindMemberById(member_id)
	//查询会员套餐
	err = productInfo.FindMemberTypeById(member_type_id)

	if !err {
		this.Abort("会员类型信息有误，请查验后再提交")
	}

	//如果是VIP会员，则直接判断
	//判断之前是否已经购买过，购买过则无需再次购买
	//订单创建流程开始
	//获取随机订单号
	orderNo = info.GetRandOrderNo()
	//订单创建
	info.OrderNo = orderNo
	info.ProductId = member_type_id
	info.ProductName = productInfo.Name
	info.MemberId = member_id

	info.FromPlatform = "pc"
	info.FromChannel = "direct"
	info.FromChannelTag = "codeshop.com"

	// info.RecommendCode = agentInfo.RecommendCode
	// info.AgentId = agentInfo.Id
	// info.AgentName = agentInfo.AgentName
	// info.AgentWeixinOpenId = agentInfo.WeixinOpenId
	// info.AgentWeixin = agentInfo.Weixin
	// info.AgentEmail = agentInfo.Email
	// info.AgentMobile = agentInfo.Mobile

	info.MemberName = minfo.Nickname
	info.MemberMobile = minfo.Mobile
	info.MemberEmail = minfo.Email
	info.MemberWeixin = minfo.Weixin
	info.MemberWeixinOpenId = minfo.WeixinOpenId

	info.CommissionAmount = 0
	info.Count = 1
	if minfo.CheckVip(member_id) {
		//VIP 会员，采用会员价购买
		info.Price = productInfo.Price
		info.Discount = 0
	} else {
		//普通会员，采用普通价格购买
		info.Price = productInfo.Price
		info.Discount = 0
	}
	info.PayAmount = info.Price
	info.Amount = info.Price
	info.IsSend = 0
	info.Status = 0
	//创建订单
	orderId, oerr := info.CreateOrder()
	if oerr == false {

		this.Abort("会员订单创建失败")
	}

	//创建微信支付记录
	payinfo.OrderId = orderId
	payinfo.OrderNo = info.OrderNo
	payinfo.PayType = 1 //消费
	payinfo.OrderType = "member"
	payinfo.MemberId = member_id
	payinfo.AgentId = info.AgentId
	payinfo.MemberName = info.MemberName
	payinfo.MemberMobile = mobile
	payinfo.MemberEmail = email
	payinfo.MemberWeixin = info.MemberWeixin
	payinfo.Amount = info.Amount
	payinfo.Discount = info.Discount
	payinfo.PayAmount = info.PayAmount
	payinfo.PayMethod = "weixin"
	payinfo.PayBody = "购买会员服务" + info.ProductName + "-优品源码网"
	payinfo.ProductId = info.ProductId
	payinfo.PayStatus = 0
	payinfo.Status = 0
	payinfo.Insert()

	//增加代理商VIP会员数
	// if agent_id > 0 && len(agent_mobile) > 0 {
	// 	agentInfo.Id = agent_id
	// 	agentInfo.AddVipMemberCount()

	// }

	url = site_pay_scan_url + "?orderno=" + orderNo
	if info.PayAmount > 0 {
		url = site_pay_scan_url + "?orderno=" + orderNo
	} else {
		//直接跳转会员详细页面
		url = fmt.Sprintf("/member/profile/%d")
	}
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")

	this.Redirect(url, 302)
	return
}

//订单支付检查
func (this *MemberHandle) Check() {
	var (
		info  = new(models.MemberOrder)
		minfo = new(models.Member)
		err   bool
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")

	order_no := strings.TrimSpace(this.GetString("orderno"))
	if order_no == "" {
		this.Abort("404")
		return
	}
	//读取数据
	err = info.FindMemberOrderByOrderNo(order_no)
	if err == false || info.Status < 1 {
		this.Abort("404")
		return
	}
	if info.IsSend == 0 && info.Status == 1 {
		//未发货状态,则进行会员增加时间处理
		err = minfo.FindMemberById(info.MemberId)
		if err == false {
			this.Abort("会员信息不存在，请联系管理员")
			return
		}
		//升级会员
		models.UpgradeVip(info.OrderNo, info.MemberId, info.ProductId)
	}
	url := fmt.Sprintf("/member/profile/")
	this.Redirect(url, 302)

}

//前台详细页
func (this *MemberHandle) Profile() {
	var (
		memberInfo *models.Member = new(models.Member)
		is_vip     string
		member_id  int64
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	token, _ := this.Ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	if IsLogin {

		err2 := memberInfo.FindMemberByIdAndToken(LoginMember.Id, token)
		if !err2 {
			member_id = 0
		} else {
			//登陆会员，则判断是否为VIP会员
			if memberInfo.CheckVip(member_id) {
				is_vip = "VIP会员"

			} else {
				is_vip = "普通会员"
			}
		}
	} else {
		this.Redirect("/login", 302)
		return
	}
	this.Data["member_id"] = LoginMember.Id
	this.Data["memberInfo"] = memberInfo
	this.Data["is_vip"] = is_vip
	this.Layout = "layout/_member_layout.html"
	this.TplName = "member/_profile.html"
}

//在线充值
func (this *MemberHandle) Recharge() {

	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		memberInfo  models.Member
		memberOrder models.MemberOrder
	)

	memberList := memberOrder.GetLastList(6)
	this.Data["PageTitle"] = "在线充值"
	this.Data["memberlist"] = memberList
	this.Data["memberInfo"] = memberInfo
	this.Layout = "layout/_member_layout.html"

	this.TplName = "member/_recharge.html"
}
