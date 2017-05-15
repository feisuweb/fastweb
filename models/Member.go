package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/sluu99/uuid"
	"time"
)

type Member struct {
	Id               int64
	Gender           int64     //性别  0  女 1  男
	Mobile           string    `orm:"unique;size(50)"`  //手机号
	Email            string    `orm:"unique;size(250)"` //邮箱
	Avatar           string    `orm:"size(500)"`        //头像
	MemberName       string    `orm:"unique;size(250)"` //会员名称
	Password         string    `orm:"size(32)"`         //密码
	PasswordSalt     string    `orm:"null;size(8)"`
	Nickname         string    `orm:"unique;size(40)"` //昵称
	Pid              int64     // 推荐用户ID  0  自己注册
	WeixinOpenId     string    `orm:"size(250)"`                   //微信开放平台ID
	Weixin           string    `orm:"size(250)"`                   //微信号码
	City             string    `orm:"size(250)"`                   //代理城市
	Province         string    `orm:"size(250)"`                   //代理城市
	Region           string    `orm:"size(250)"`                   //代理区域
	Address          string    `orm:"size(500)"`                   //代理地址
	LastLoginTime    time.Time `orm:"auto_now_add;type(datetime)"` //最后登录时间
	RegisterTime     time.Time `orm:"auto_now_add;type(datetime)"` //注册时间
	VipExpire        time.Time `orm:"auto_now_add;type(datetime)"` //vip过期时间
	MemberType       int64     // 会员类型ID 1 一年VIP会员 2 终身VIP会员
	LoginTimes       int64     //登录次数
	LastLoginIp      string    `orm:"size(32)"` //最后登录IP
	RegisterIp       string    `orm:"size(32)"` //第一次注册时候的IP
	IsVip            int64     //是否为VIP
	IsValidateMobile int64     //是否验证手机号
	IsValidateEmail  int64     //是否邮箱地址
	Points           int64     //用户积分
	Money            int64     //金钱数量
	VipLevel         int64     //VIP等级
	RecommendCode    string    //推荐码
	AgentMobile      string    //代理商手机号
	AgentId          int64     //代理商会员ID
	Token            string    //token
	AddTime          time.Time `orm:"auto_now_add;type(datetime)"`
	MemberActivated  int64     `orm:"default(0);size(2)"` //1 激活 0 未激活
	Status           int64     // 0  正常  -1 封号  1 限制登录
}

func FindMemberByToken(token string) (bool, Member) {
	o := orm.NewOrm()
	var memberInfo Member
	err := o.QueryTable(memberInfo).Filter("Token", token).One(&memberInfo)
	return err != orm.ErrNoRows, memberInfo
}

func GetMemberNameById(memberId int64) string {
	info := new(Member)
	err := info.FindMemberById(memberId)

	if err {
		return info.MemberName
	} else {
		return ""
	}

}

func (m *Member) ChangePassword(id int64, oldPassword string, newPassword string) bool {
	var (
		pwd  string
		pwd2 string
	)
	err := m.FindMemberById(id)
	if err {
		salt := m.PasswordSalt

		pwd = Md5(oldPassword + salt)
		if pwd == m.Password { //如果老密码正确，则修改新密码
			pwd2 = Md5(newPassword + salt)
			m.Password = pwd2
			m.Update("password")
			return true
		} else {
			return false
		}

	}
	return false
}

func (m *Member) Register() bool {

	o := orm.NewOrm()
	var pwd string
	var token = uuid.Rand().Hex()
	salt := GetRandomSalt()
	pwd = Md5(m.Password + salt)
	m.PasswordSalt = salt
	m.Password = pwd
	m.Token = token
	m.Avatar = "/static/img/avatar_default.png"
	_, err := o.Insert(m)
	return err != orm.ErrNoRows
}

func (m *Member) Login(username string, password string, ip string) bool {
	var pwd string
	err := m.FindMemberByMemberName(username)
	if err {

		salt := m.PasswordSalt
		pwd = Md5(password + salt)
		if m.Password == pwd {

			if len(m.Token) < 8 {
				var token = uuid.Rand().Hex()
				m.Token = token
				//更新登录信息
				m.LastLoginIp = ip
				m.LastLoginTime = time.Now()
				m.Update("Token", "LastLoginIp", "LastLoginTime")
			} else {
				//更新登录信息
				m.LastLoginIp = ip
				m.LastLoginTime = time.Now()
				m.Update("LastLoginIp", "LastLoginTime")
			}

			//记录登录日志
			return true
		} else {
			return false
		}

	}

	return false
}

func (m *Member) CheckVip(memberId int64) bool {
	o := orm.NewOrm()
	var mem Member
	err := o.QueryTable(mem).Filter("Id", memberId).Filter("IsVip", 1).One(&mem)
	if err != orm.ErrNoRows {
		return false
	}
	//判断是否过期会员
	t := time.Now()
	//判断会员是否过期
	ret := t.Before(mem.VipExpire)
	if ret {
		return true
	}
	return false

}

func (m *Member) FindMemberByIdAndToken(memberId int64, token string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", memberId).Filter("Token", token).One(m)
	return err != orm.ErrNoRows

}

func (m *Member) FindMemberByMemberName(username string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("MemberName", username).One(m)
	return err != orm.ErrNoRows
}

func (m *Member) FindMemberById(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}

func (m *Member) FindMemberByMobileAndEmail(mobile string, email string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Email", email).Filter("Mobile", mobile).One(m)
	return err != orm.ErrNoRows
}

func (m *Member) FindMemberByMobileOrEmail(mobile string, email string) bool {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("Email", email).Or("Mobile", mobile)
	qs := o.QueryTable(m)
	qs = qs.SetCond(cond1)
	err := qs.One(m)
	return err != orm.ErrNoRows
}

func (m *Member) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Member) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Member) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Member) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Member) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

///最新列表
func (m *Member) GetLastList(pagesize int64) []*Member {
	var info Member
	list := make([]*Member, 0)

	info.Query().OrderBy("-id").Limit(pagesize, 0).All(&list, "Id", "MemberName", "Mobile", "MemberType", "AddTime")

	return list
}
