package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MemberSafeLog struct {
	Id      int64     //id
	Action  string    `orm:"size(100)"` //动作 changepassword，login，logout，buy，download，view
	Content string    `orm:"type(text)"`
	Member  *Member   `orm:"rel(fk)"` //会员
	Ip      string    `orm:"size(32)"`
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *MemberSafeLog) AddSafeLog(member Member, action string, ip string, content string) error {
	var msli MemberSafeLog
	msli.Member = &member
	msli.Action = action
	msli.Ip = ip
	msli.Content = content
	if _, err := orm.NewOrm().Insert(&msli); err != nil {
		return err
	}
	return nil
}

func (m *MemberSafeLog) FindMemberSafeLogByMemberId(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}

func (m *MemberSafeLog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MemberSafeLog) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MemberSafeLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MemberSafeLog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *MemberSafeLog) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
