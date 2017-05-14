package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/sluu99/uuid"
	"time"
)

type User struct {
	Id            int64
	Username      string    `orm:"unique;size(50)"`
	Password      string    `orm:"size(32)"`
	Nickname      string    `orm:"unique;size(50)"`
	PasswordSalt  string    `orm:"null;size(8)"`
	Avatar        string    `orm:"null"`
	Email         string    `orm:"unique;null"`
	Mobile        string    `orm:"unique;null"`
	RegisterIp    string    `orm:"null"`
	Signature     string    `orm:"size(500)"`
	RegisterTime  time.Time `orm:"auto_now_add;type(datetime)"`
	LastLoginTime time.Time `orm:"auto_now_add;type(datetime)"`
	LoginTimes    int64     `orm:"default(0)"`
	LastLoginIp   string    `orm:"size(32)"`
	Token         string    `orm:"unique"`
	AddTime       time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	UpdateTime    time.Time `orm:"auto_now_add;type(datetime)"` //更新某季下载的时间，用来排序最近更新
	Status        int64     `orm:"default(0)"`                  //状态:0为显示 -1为不显示
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
func GetUserById(id int64) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}
func FindUserById(id int64) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func Register(user *User) int64 {

	o := orm.NewOrm()

	flag, _ := FindUserByUserName(user.Username)

	if flag {
		return 1
	}
	var token = uuid.Rand().Hex()
	var pwd string
	salt := GetRandomSalt()
	pwd = GetEncryptPassword(user.Password, salt)
	user.PasswordSalt = salt
	user.Password = pwd
	user.Token = token
	_, err := o.Insert(user)
	if err != orm.ErrNoRows {
		return 0
	} else {
		return -1
	}

}

func Login(username string, password string, ip string) (bool, User) {
	var (
		pwd        string
		user       User
		userUpdate User
		err        bool
	)
	err, user = FindUserByUserName(username)
	if !err {
		return false, user
	}

	pwd = GetEncryptPassword(password, user.PasswordSalt)
	o := orm.NewOrm()

	err2 := o.QueryTable(user).Filter("Username", username).Filter("Password", pwd).One(&user)
	if err2 != orm.ErrNoRows {
		//更新登录时间和IP
		o := orm.NewOrm()
		userUpdate.LoginTimes += 1
		userUpdate.Id = user.Id
		userUpdate.LastLoginIp = ip
		userUpdate.LastLoginTime = time.Now()
		o.Update(&userUpdate, "LastLoginTime", "LastLoginIp")
		return true, user
	} else {
		return false, user
	}

}
