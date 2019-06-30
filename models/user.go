package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int       `orm:"column(id);auto" description:"id"`
	Username string    `orm:"column(username);unique" description:"用户名"`
	Password string    `orm:"column(password)" description:"密码"`
	Login    time.Time `orm:column(latest_login);auto_now" description:"上次登陆时间"`
	Profile
	Project *Project `orm:"reverse(one)"`
	BaseModel
}

type Profile struct {
	Gender  string `orm:"column(gender);default('')" description:"性别"`
	Age     int    `orm:"column(age);default(0)" description:"年龄"`
	Address string `orm:"column(address);default('')" description:"地址"`
	Email   string `orm:"column(email);default('')" description:"邮箱"`
}

// 注册用户
func AddUser(u *User) (int64, error) {
	return orm.NewOrm().Insert(u)
}

func GetUserByName(username string) (*User, error) {
	u := new(User)
	o := orm.NewOrm()
	err := o.QueryTable(TableName("user")).Filter("username", username).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetAllUsers() ([]User, error) {
	var u []User
	_, err := orm.NewOrm().QueryTable(TableName("user")).All(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UpdateUser(user *User, fields ...string) error {
	_, err := orm.NewOrm().Update(user, fields...)
	return err
}
