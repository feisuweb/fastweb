package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//工单处理记录表
type IssuesLog struct {
	Id                int64
	IssueId           int64     //工单ID
	MemberId          int64     // 发起会员ID
	ProcessUserId     int64     // 工单最后受理UserID
	IssueContent      string    `orm:"type(text)"`                  //工单内容
	IssueReplyContent string    `orm:"type(text)"`                  //最后反馈内容
	IssueReplyTime    time.Time `orm:"auto_now_add;type(datetime)"` //工单最后回复时间
	AddTime           time.Time `orm:"auto_now_add;type(datetime)"` //工单最开始发起时间
	Status            int64     // 0 未回复 1 已回复
}

func (m *IssuesLog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *IssuesLog) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *IssuesLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *IssuesLog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *IssuesLog) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
