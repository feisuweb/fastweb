package notify

import (
	"net/smtp"
	"strings"
)

//==================邮件发送基础封装=====================

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	to = strings.Trim(to, " ")
	if len(to) == 0 {
		return nil
	}
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendMailNotify(to string, subject string, content string) error {
	user := "service@codeshop.com"
	password := "T5hJMMcy"
	host := "smtp.mxhichina.com:25"
	err := SendToMail(user, password, host, to, subject, content, "html")
	return err
}

//================================客户通知==============================
//发送给客户邮件-产品订单通知
func SendToCustomerMailOrderNotify(m *NotifyInfo) error {

	if len(m.MemberEmail) > 0 {
		return nil
	}

	subject := "$SiteName$-订单号:$OrderNo$ ,产品下载地址"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>产品下载地址
        </h3>
        <br/>
        恭喜 ，您在$SiteName$的订单 $OrderNo$ 支付成功，以下是产品下载地址：<br/>
        <b>产品名称：$ProductName$</b><br/>
        <b>产品本地下载地址是：$DownloadUrl$</b><br/>
        <b>产品网盘下载地址是：$YunpanDownloadUrl$</b><br/>
        <b>产品提取码是：$DownloadCode$</b><br/>
        <b>产品解压密码是：$UnzipPassword$</b><br/>

        </body>
        </html>
        `

	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(m.MemberEmail, subject, body)
	return err
}

//发送给客户邮箱-会员订单通知
func SendToCustomerMailMemberOrderNotify(m *NotifyInfo) error {

	if len(m.MemberEmail) > 0 {
		return nil
	}
	subject := "$SiteName$-恭喜您，购买会员成功！"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>产品下载地址
        </h3>
        <br/>
        恭喜 ，您在$SiteName$的购买会员订单 $OrderNo$ 支付成功，以下是会员信息：<br/>
        <b>会员名称：$MemberName$</b><br/>
        <b>手机号：$MemberMobile$</b><br/>
        <b>邮箱：$MemberEmail$</b><br/>
        <b>登陆密码：$Password$</b><br/>
        <b>会员类型：$MemberTypeName$</b><br/>
        您可以登录网站 codeshop.com 下载优秀的产品资源。
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(m.MemberEmail, subject, body)
	return err
}

//================================卖家通知==============================
//发送给卖家邮件-产品订单通知
func SendToSellerMailOrderNotify(m *NotifyInfo) error {
	if len(m.SellerEmail) > 0 {
		return nil
	}
	subject := "$SiteName$-卖出产品-订单号:$OrderNo$-下单时间:$AddTime$"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>卖出一份产品订单号:$OrderNo$
        </h3>
        <br/>
        恭喜 ，客户在$SiteName$下单成功，以下是订单信息：<br/>
        <b>订单号是：$OrderNo$</b><br/>
        <b>产品名称：$ProductName$</b><br/>
        <b>订单金额：$Amount$</b><br/>
        <b>支付方式：$PayMethod$</b><br/>
        <b>客户手机：$MemberMobile$</b><br/>
        <b>客户邮箱：$MemberEmail$</b><br/>
        <b>客户微信：$MemberWeixin$</b><br/>
        <b>下单时间：$AddTime$</b><br/>
        <b>付款时间：$PayTime$</b><br/>
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(m.SellerEmail, subject, body)
	return err
}

//====================站长通知=====================
//发送给网站站长邮件-会员订单通知
func SendToMasterMailMemberOrderNotify(m *NotifyInfo) error {

	subject := "$SiteName$-卖出VIP会员-订单号:$OrderNo$-下单时间:$AddTime$"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>VIP会员销售通知
        </h3>
        <br/>
        恭喜站长,您的网站$SiteName$成功卖出一份VIP会员套餐：<br/>
        <b>订单号码：$OrderNo$</b><br/>
        <b>套餐名称：$ProductName$</b><br/>
        <b>订单金额：$Amount$</b><br/>
        <b>会员名称：$MemberName$</b><br/>
        <b>会员手机：$MemberMobile$</b><br/>
        <b>会员邮箱：$MemberEmail$</b><br/>
        <b>会员微信：$MemberWeixin$</b><br/>
        
        <b>购买日期：$PayTime$</b><br/>
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(MasterEmail, subject, body)
	return err
}

//================================代理商通知==============================
//发送给代理商邮件-产品订单通知
func SendToAgentMailOrderNotify(m *NotifyInfo) error {

	if len(m.AgentEmail) > 0 {
		return nil
	}

	subject := "$SiteName$-您推荐的产品已卖出-佣金金额:$CommissionAmount$-下单时间:$AddTime$"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>产品下载地址
        </h3>
        <br/>
        恭喜 ，您推荐的$SiteName$产品$ProductName$被会员$MemberName$购买,并支付成功，以下是订单信息：<br/>
        <b>订单号是：$OrderNo$</b><br/>
        <b>产品名称：$ProductName$</b><br/>
        <b>订单金额：$Amount$</b><br/>
        <b>佣金金额：$CommissionAmount$</b><br/>
        <b>支付方式：$PayMethod$</b><br/>
        <b>下单时间：$AddTime$</b><br/>
        <b>付款时间：$PayTime$</b><br/>
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(m.AgentEmail, subject, body)
	return err
}

//发送给代理商邮件-会员订单通知
func SendToAgentMailMemberOrderNotify(m *NotifyInfo) error {

	subject := "$SiteName$-恭喜代理商:$MemberName$ 购买VIP套餐一份,佣金金额:$Amount$-下单时间:$AddTime$"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>代理商订单信息
        </h3>
        <br/>
        恭喜 ，您推荐的$SiteName$会员$MemberName$购买了VIP会员服务套餐$ProductName$,以下是订单信息：<br/>
        <b>订单号码：$OrderNo$</b><br/>
        <b>套餐名称：$ProductName$</b><br/>
        <b>订单金额：$Amount$</b><br/>
        <b>会员名称：$MemberName$</b><br/>
        <b>佣金金额：$CommissionAmount$</b><br/>
        <b>下单时间：$AddTime$</b><br/>
        <b>付款时间：$PayTime$</b><br/>
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)

	err := SendMailNotify(m.AgentEmail, subject, body)
	return err
}

//===================密码修改通知==============
//发送给会员密码修改通知
func SendToMemberMailPasswordChangedNotify(m *NotifyInfo) error {

	subject := "$SiteName$-密码修改通知，会员 $MemberName$ 在 $ChangePasswordTime$ 修改了密码。"

	body := `
        <html>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <body>
        <h3>密码修改通知
        </h3>
        <br/>
        您好 ，您在 $SiteName$ 的会员账号 $MemberName$ 于 $ChangePasswordTime$ 修改了密码：<br/>
        <b>新的密码是：$ChangePasswordNewPassword$</b><br/>
        <b>修改IP地址是：$ChangePasswordIp$</b><br/>
        如果不是您本人修改，请及时通过找回密码找回，或者点击冻结账号链接立刻冻结账号(24小时内有效)。  <a href="http://www.codeshop.com/member/lock?id=$MemberSafeLogId$" target="_blank">冻结账号</a>
        </body>
        </html>
        `
	subject = ReplaceNotifyContent(m, subject)
	body = ReplaceNotifyContent(m, body)
	err := SendMailNotify(m.MemberEmail, subject, body)
	return err
}
