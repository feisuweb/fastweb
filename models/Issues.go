package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//工单表
type Issues struct {
	Id                int64
	ProductId         int64     // 产品ID  默认0
	IssueType         int64     // 工单类型 0  产品问题 1  服务问题 2 会员问题
	MemberId          int64     // 发起会员ID
	ProcessUserId     int64     // 工单最后受理UserID
	IssueTitle        string    // 工单问题标题
	IssueContent      string    `orm:"type(text)"`                  //工单内容
	IssueOpenTime     time.Time `orm:"auto_now_add;type(datetime)"` //开启时间
	IssueCloseTime    time.Time `orm:"auto_now_add;type(datetime)"` //关闭时间
	IssueStatus       int64     // 0  未支付  1 等待支付回调 2 支付成功 3 支付失败
	IssueReplyContent string    `orm:"type(text)"`                  //最后反馈内容
	IssueReplyTime    time.Time `orm:"auto_now_add;type(datetime)"` //工单最后回复时间
	IssueProcessTime  time.Time `orm:"auto_now_add;type(datetime)"` //工单处理时间
	Satisfaction      int64     //工单处理满意度 0-5 默认5
	AddTime           time.Time `orm:"auto_now_add;type(datetime)"` //工单最开始发起时间
	Status            int64     // 0 等待受理 1 工单受理成功  2 工单关闭
}

func (m *Issues) FindPayByIssueNo(issueNo string) error {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("IssueNo", issueNo).One(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Issues) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Issues) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Issues) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Issues) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Issues) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
