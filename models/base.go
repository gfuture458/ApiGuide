package models

import "time"

type BaseModel struct {
	CreateAt time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" description:"修改时间"`
	IsActive bool      `orm:"default(1);type(bool)" description:"是否可用"`
}

func (u *User) TableName() string {
	return TableName("user")
}
