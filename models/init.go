package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/feisuweb/fastweb/libs/notify"
	"time"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")

	orm.RegisterModelWithPrefix("fastweb_", new(Member))
	orm.RegisterModelWithPrefix("fastweb_", new(MemberOrder))
	orm.RegisterModelWithPrefix("fastweb_", new(MemberType))
	orm.RegisterModelWithPrefix("fastweb_", new(User))
	orm.RegisterModelWithPrefix("fastweb_", new(PayLog))
	orm.RegisterModelWithPrefix("fastweb_", new(MemberSafeLog))

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8&loc=Asia%2FShanghai")

	name := "default" //数据库别名
	force := false    //不强制建数据库
	verbose := true   //打印建表过程
	orm.RunSyncdb(name, force, verbose)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//管理员
	initAdmin()
}

func initAdmin() {
	//数据初始化
	var (
		flag bool
	)
	//管理员初始化

	flag, _ = FindUserByUserName("admin")
	if !flag {

		//如果没有超高级管理员，则初始化一个。
		var user User
		user.Id = 1
		user.Username = "admin"
		user.Password = "123456"
		user.Nickname = "超级管理员"
		user.Avatar = "/static/imgs/avatar.png"
		user.Signature = "这家伙很懒，什么都没留下~"
		user.RegisterIp = "127.0.0.1"
		Register(&user)
	}
}

//升级会员
func UpgradeVip(orderNo string, memberId int64, memberTypeId int64) bool {
	//升级会员期限
	//如果之前是VIP会员，则进行累加
	//会员期限=已有的期限+购买的期限
	//如果之前不是会员则直接进行更新期限
	var minfo Member
	var moinfo MemberOrder
	var mtinfo MemberType
	var notifyInfo notify.NotifyInfo

	moinfo.FindMemberOrderByOrderNo(orderNo)
	minfo.FindMemberById(memberId)
	mtinfo.FindMemberTypeById(memberTypeId)
	if moinfo.Status < 1 || moinfo.Status == 2 {
		return false
	}
	//判断是否过期会员
	t := time.Now()
	//判断会员是否过期
	ret := t.Before(minfo.VipExpire)
	if ret {
		//会员过期,在当前基础上增加时间
		t1 := time.Duration(mtinfo.ValidTime) * 24 * time.Hour
		t.Add(t1)
		minfo.VipExpire = t

	} else {
		//会员没有过期，则在这个基础上增加时间
		t2 := time.Duration(mtinfo.ValidTime) * 24 * time.Hour
		minfo.VipExpire.Add(t2)

	}
	//设置会员为VIP
	minfo.IsVip = 1
	minfo.Update()
	//更新会员订单表
	moinfo.IsSend = 1
	moinfo.Status = 2
	moinfo.Update()
	//通知

	//通知信息赋值
	//订单信息
	notifyInfo.OrderNo = orderNo
	notifyInfo.Amount = moinfo.Amount
	notifyInfo.PayMethod = moinfo.PayMethod

	//产品信息
	notifyInfo.ProductId = moinfo.ProductId
	notifyInfo.ProductName = moinfo.ProductName

	//会员信息
	notifyInfo.MemberId = memberId
	notifyInfo.MemberEmail = moinfo.MemberEmail
	notifyInfo.MemberMobile = moinfo.MemberMobile
	notifyInfo.MemberName = moinfo.MemberName

	//推荐者信息
	notifyInfo.AgentId = moinfo.AgentId
	notifyInfo.AgentName = moinfo.AgentName
	notifyInfo.AgentEmail = moinfo.AgentEmail
	notifyInfo.AgentMobile = moinfo.AgentMobile
	notifyInfo.AgentWeixinOpenId = moinfo.AgentWeixinOpenId

	//给客户发送会员订单通知
	notify.SendToCustomerMemberOrderNotify(&notifyInfo)
	//给站长发送会员卖出通知
	notify.SendToMasterMemberOrderNotify(&notifyInfo)
	//给推荐者发送会员卖出通知
	notify.SendToAgentMemberOrderNotify(&notifyInfo)
	return true
}
