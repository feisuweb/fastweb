package notify

import (
	"fmt"
	"strings"
)

var (
	MasterEmail        string = "guoxinzz@163.com"
	MasterMobile       string = "13926485656"
	MasterWeixinOpenId string = "wx1234567890"
	SiteName           string = "优品源码网"
	WebSiteUrl         string = "http://www.baidu.com"
)

//通知表
type NotifyInfo struct {
	NotifyId           int64
	OrderNo            string //订单号码
	MemberId           int64  //会员ID
	MemberName         string //会员名称
	MemberEmail        string //会员邮箱
	MemberMobile       string //会员手机号
	MemberWeixin       string //代理商微信号
	MemberQQ           string //代理商QQ号
	MemberWeixinOpenId string //会员微信OPENID
	IsVip              bool   //是否为VIP会员
	//密码修改通知内容
	ChangePasswordTime        string //修改时间
	ChangePasswordIp          string //ip地址
	ChangePasswordNewPassword string //新密码
	MemberSafeLogId           string //日志记录ID，用来锁定账号等信息

	AgentId     int64
	AgentName   string //代理商名称
	AgentEmail  string //代理商邮箱
	AgentMobile string //代理商手机号

	AgentWeixinOpenId string //代理商微信OPENID

	SellerId           int64
	SellerName         string //卖家名称
	SellerEmail        string //卖家邮箱
	SellerMobile       string //卖家手机号
	SellerWeixinOpenId string //卖家微信OPENID

	ProductId        int64   //产品ID
	ProductName      string  //产品名字
	Amount           float64 //订单金额
	CommissionAmount float64 //佣金金额
	PayMethod        string  //支付方式

	DownloadUrl        string //本地下载地址
	DownloadYunPanUrl  string //云盘地址
	DownloadYunPanCode string //云盘提取码
	UnzipPassword      string //解压密码

	AddTime string //订单创建时间 内容默认从time 转换成string
	PayTime string //订单付款时间  内容默认从time 转换成string
}

//=====================通知基础===================

//模板内容替换
func ReplaceNotifyContent(m *NotifyInfo, content string) string {

	//站点信息替换
	content = strings.Replace(content, "$SiteName$", SiteName, -1)

	amount := fmt.Sprintf("%.2f元", m.Amount)
	commissionAmount := fmt.Sprintf("%.2f元", m.CommissionAmount)
	payMethod := "微信扫码"
	if m.PayMethod == "weixinscan" {
		payMethod = "微信扫码"
	} else if m.PayMethod == "weixinwap" {
		payMethod = "微信支付"
	} else if m.PayMethod == "alipay" {
		payMethod = "支付宝"
	} else {
		payMethod = "微信扫码"
	}
	//替换内容
	content = strings.Replace(content, "$SiteName$", "优品源码网", -1) //网站名称
	//会员信息
	content = strings.Replace(content, "$MemberName$", m.MemberName, -1)           //会员名称
	content = strings.Replace(content, "$OrderNo$", m.OrderNo, -1)                 //订单号
	content = strings.Replace(content, "$ProductName$", m.ProductName, -1)         //产品名称
	content = strings.Replace(content, "$Amount$", amount, -1)                     //订单金额
	content = strings.Replace(content, "$CommissionAmount$", commissionAmount, -1) //佣金金额
	content = strings.Replace(content, "$PayMethod$", m.PayMethod, -1)             //支付方式

	content = strings.Replace(content, "$MemberName$", m.MemberName, -1)     //会员名称
	content = strings.Replace(content, "$MemberEmail$", m.MemberEmail, -1)   //会员邮箱
	content = strings.Replace(content, "$MemberMobile$", m.MemberMobile, -1) //会员手机号
	content = strings.Replace(content, "$MemberWeixin$", m.MemberWeixin, -1) //会员微信号
	content = strings.Replace(content, "$MemberQQ$", m.MemberQQ, -1)         //会员QQ号

	content = strings.Replace(content, "$OrderNo$", m.OrderNo, -1)                 //订单号
	content = strings.Replace(content, "$ProductName$", m.ProductName, -1)         //产品名称
	content = strings.Replace(content, "$Amount$", amount, -1)                     //订单金额
	content = strings.Replace(content, "$CommissionAmount$", commissionAmount, -1) //佣金金额
	content = strings.Replace(content, "$PayMethod$", payMethod, -1)               //支付方式

	content = strings.Replace(content, "$AgentName$", m.AgentName, -1)     //代理商名称
	content = strings.Replace(content, "$AgentEmail$", m.AgentEmail, -1)   //代理商邮箱
	content = strings.Replace(content, "$AgentMobile$", m.AgentMobile, -1) //代理商手机号

	content = strings.Replace(content, "$AddTime$", m.AddTime, -1) //下单时间
	content = strings.Replace(content, "$PayTime$", m.PayTime, -1) //支付时间

	//下载地址替换
	content = strings.Replace(content, "$DownloadUrl$", m.DownloadUrl, -1)
	content = strings.Replace(content, "$YunpanDownloadUrl$", m.DownloadYunPanUrl, -1)
	content = strings.Replace(content, "$DownloadCode$", m.DownloadYunPanCode, -1)
	content = strings.Replace(content, "$UnzipPassword$", m.UnzipPassword, -1)

	content = strings.Replace(content, "$AddTime$", m.AddTime, -1) //下单时间
	content = strings.Replace(content, "$PayTime$", m.PayTime, -1) //支付时间

	//修改密码内容替换
	content = strings.Replace(content, "$ChangePasswordTime$", m.ChangePasswordTime, -1)               //修改密码时间
	content = strings.Replace(content, "$ChangePasswordNewPassword$", m.ChangePasswordNewPassword, -1) //修改密码的新密码
	content = strings.Replace(content, "$ChangePasswordIp$", m.ChangePasswordIp, -1)                   //修改密码的IP地址
	content = strings.Replace(content, "$MemberSafeLogId$", m.MemberSafeLogId, -1)                     //会员安全日志ID

	return content
}

//=========================产品订单通知====================
//发送客户订单通知
func SendToCustomerOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToCustomerMailOrderNotify(m)
	//短信通知
	SendToCustomerSMSOrderNotify(m)
	//微信通知
	SendToCustomerWeixinOrderNotify(m)
}

//发送卖家订单通知
func SendToSellerOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToSellerMailOrderNotify(m)
	//短信通知
	SendToSellerSMSOrderNotify(m)
	//微信通知
	SendToSellerWeixinOrderNotify(m)

}

//发送代理商订单通知
func SendToAgentOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToAgentMailOrderNotify(m)
	//短信通知
	SendToAgentSMSOrderNotify(m)
	//微信通知
	SendToAgentWeixinOrderNotify(m)

}

//==================会员订单通知===========================
//发送客户会员订单通知
func SendToCustomerMemberOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToCustomerMailMemberOrderNotify(m)
	//短信通知
	SendToCustomerSMSMemberOrderNotify(m)
	//微信通知
	SendToCustomerWeixinMemberOrderNotify(m)
}

//发送给站长会员订单通知
func SendToMasterMemberOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToMasterMailMemberOrderNotify(m)
	//短信通知
	SendToMasterSMSMemberOrderNotify(m)
	//微信通知
	SendToMasterWeixinMemberOrderNotify(m)

}

//发送代理商会员订单通知
func SendToAgentMemberOrderNotify(m *NotifyInfo) {
	//邮件通知
	SendToAgentMailMemberOrderNotify(m)
	//短信通知
	SendToAgentSMSMemberOrderNotify(m)
	//微信通知
	SendToAgentWeixinMemberOrderNotify(m)

}

//===================密码修改通知==============
//发送给会员密码被修改的通知
func SendToMemberPasswordChangedNotify(m *NotifyInfo) {
	//邮件通知
	SendToMemberMailPasswordChangedNotify(m)
	//短信通知
	SendToMemberSMSPasswordChangedNotify(m)
	//微信通知
	SendToMemberWeixinPasswordChangedNotify(m)
}
