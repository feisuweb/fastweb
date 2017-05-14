package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//账单记录表
type PayLog struct {
	Id               int64
	OrderId          int64
	OrderNo          string `orm:"size(100)"` //交易号
	ProductId        int64  //产品ID
	OrderType        string
	PayType          int64     //  交易类型 0  充值  1  消费  2 赠送
	MemberId         int64     //会员ID
	AgentId          int64     //代理商ID
	MemberName       string    `orm:"size(150)"`
	MemberMobile     string    `orm:"size(150)"`
	MemberEmail      string    `orm:"size(500)"`
	MemberWeixin     string    `orm:"size(500)"`
	Amount           float64   //交易金额
	Discount         float64   //折扣金额
	PayAmount        float64   //支付金额  =交易金额  - 抵扣
	PayMethod        string    `orm:"size(100)"`                   //alipay (支付宝支付)  weixin(微信支付) agent （代理商支付）
	PayBody          string    `orm:"type(text)"`                  //交易内容
	PayTime          time.Time `orm:"auto_now_add;type(datetime)"` //支付时间
	PayStatus        int64     // 0  未支付  1 等待支付回调 2 支付成功 3 支付失败
	PayNotifyContent string    `orm:"type(text)"`                  //支付通知接口反馈内容
	PayNotifyTime    time.Time `orm:"auto_now_add;type(datetime)"` //交易时间
	AddTime          time.Time `orm:"auto_now_add;type(datetime)"` //交易时间
	Status           int64     // 0 等待支付  1 交易成功  2 交易失败
}

func (m *PayLog) FindPayByOrderNo(orderNo string) error {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("OrderNo", orderNo).One(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *PayLog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *PayLog) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *PayLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *PayLog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *PayLog) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
