package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//加密密码
func GetEncryptPassword(password string, salt string) string {
	//GetRandomSalt()
	var pwd string
	pwd = Md5(Md5(password + salt))
	return pwd
}

func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

func ValidateMobile(mobile string) bool {
	/*
		目前移动、联通、电信三大运营商的手机号段大致如下：
		1、移动号段有134,135,136,137,138,139,147,150,151,152,157,158,159,178,182,183,184,187,188。
		2、联通号段有130，131，132，155，156，185，186，145，176。
		3、电信号段有133，153，177，180，181，189。
	*/
	reg := regexp.MustCompile("^(13[0-9]|14[57]|15[0-35-9]|18[0-9]|17[07-9])\\d{8}$")
	return reg.MatchString(mobile)
}

func ValidateEmail(email string) bool {
	reg := regexp.MustCompile(`^([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})$`)
	return reg.MatchString(email)
}

func ValidateLength(value string, min, max int) bool {
	l := len(value)
	if min > 0 && l < min {
		return false
	}
	if max > 0 && l > max {
		return false
	}
	return true
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(l int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ReplaceMobile(mobile string) string {
	m1 := mobile[0:3]
	m2 := mobile[7:11]
	return fmt.Sprintf("%s****%s", m1, m2)
}
