package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//会员类型模型
type MemberType struct {
	Id         int64
	Name       string    `orm:"size(500)"`  //会员类型中文名称
	Desc       string    `orm:"size(1000)"` //会员描述
	ValidTime  int64     //会员类型按天计算有效时间
	Price      float64   //购买价格
	AddTime    time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64     //会员类型:0为有效 -1为无效 1为推荐
}

func (m *MemberType) FindProductById(productId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", productId).One(m)
	if err != nil {
		return false
	}
	return true
}

func GetMemberTypeNameById(memberTypeId int64) string {
	info := new(MemberType)
	err := info.FindMemberTypeById(memberTypeId)
	if err {
		return info.Name
	} else {
		return "普通会员"
	}

}
func (m *MemberType) GetMemberTypeNameById(memberTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", memberTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *MemberType) FindMemberTypeById(memberTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", memberTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *MemberType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *MemberType) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MemberType) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MemberType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MemberType) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
