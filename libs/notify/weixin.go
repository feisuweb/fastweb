package notify

import (
	"github.com/astaxie/beego"
	"strings"
)

//send weixin msg
func SendWeixinNotify(weixinOpenId string, content string) error {

	weixinOpenId = strings.Trim(weixinOpenId, " ")
	if len(weixinOpenId) == 0 {
		return nil
	}
	beego.Info(weixinOpenId + "----weixin----" + content)
	return nil
}

//================================客户通知==============================
//发送给客户短信-产品订单通知
func SendToCustomerWeixinOrderNotify(m *NotifyInfo) error {
	body := `恭喜，您的订单$OrderNo$已经发货,发货信息:网盘下载地址：$YunpanDownloadUrl$ 提取码：$DownloadCode$ 解压密码：$UnzipPassword$  本地下载地址：$DownloadUrl$`
	body = ReplaceNotifyContent(m, body)
	err := SendWeixinNotify(m.MemberWeixinOpenId, body)
	return err
}

//发送给客户邮箱-会员订单通知
func SendToCustomerWeixinMemberOrderNotify(m *NotifyInfo) error {
	body := `恭喜您,您成功升级为 $ProductName$ ,您的会员账号是 $MemberName$ 默认密码:$Password$ 手机号:$MemberMobile$ 邮箱 $MemberEmail$ ,登录codeshop.com 优惠下载更多商业产品 `
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.MemberWeixinOpenId, body)
	return err
}

//================================卖家通知==============================
//发送给卖家短信-产品订单通知
func SendToSellerWeixinOrderNotify(m *NotifyInfo) error {

	body := `卖出产品,下单客户:$MemberName$，订单号:$OrderNo$ 金额是：$Amount$ 名称：$ProductName$ 客户手机：$MemberMobile$，邮箱: $MemberEmail$  下单时间:$AddTime$  `

	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.SellerWeixinOpenId, body)
	return err
}

//===========================发送给站长会员购买通知=======================
//发送给站长短信-会员订单通知
func SendToMasterWeixinMemberOrderNotify(m *NotifyInfo) error {

	body := `卖出会员,会员$MemberName$购买VIP套餐$ProductName$，订单号:$OrderNo$ 订单金额是：$Amount$ 代理商佣金：$CommissionAmount$  代理商手机:$AgentMobile$  会员手机：$MemberMobile$，会员邮箱: $MemberEmail$ `

	body = ReplaceNotifyContent(m, body)
	err := SendWeixinNotify(MasterWeixinOpenId, body)
	return err
}

//================================代理商通知==============================
//发送给卖家短信-产品订单通知
func SendToAgentWeixinOrderNotify(m *NotifyInfo) error {

	body := `推荐产品$ProductName$被客户$MemberName$ 购买,获取佣金:$CommissionAmount$.`

	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.AgentWeixinOpenId, body)
	return err
}

//发送给代理商短信-会员订单通知
func SendToAgentWeixinMemberOrderNotify(m *NotifyInfo) error {
	body := `代理商$MemberName$ 在$AddTime$ 购买VIP套餐$ProductName$，获得佣金额是：$CommissionAmount$`
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.AgentWeixinOpenId, body)
	return err
}

//===================密码修改通知==============
//发送给代理商短信-会员订单通知
func SendToMemberWeixinPasswordChangedNotify(m *NotifyInfo) error {
	body := `尊敬的VIP会员，您在$SiteName$的账号 $MemberName$ 于$PassowrdChangeTime$ 修改了密码，新密码为$ChangePasswordNewPassword$ 。如果不是您本人修改，请及时通过找回密码找回，或者点击冻结账号链接立刻冻结账号(24小时内有效)。  <a href="http://www.codeshop.com/member/lock?id=$MemberSafeLogId$">冻结账号</a>`
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.MemberWeixinOpenId, body)
	return err
}
