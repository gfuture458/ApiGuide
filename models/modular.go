package models

import (
	"github.com/astaxie/beego/orm"
)

type Modular struct {
	Id        int          `orm:"column(id);auto" description:"id" json:"id"`
	Name      string       `orm:"column(name);unique" description:"模块名称" json:"name"`
	Prefix    string       `orm:"column(prefix)" description:"路由前缀" json:"prefix"`
	Project   *Project     `orm:"null;rel(fk);on_delete(set_null)" description:"关联项目" json:"project"`
	Interface []*Interface `orm:"reverse(many)"`
	BaseModel
}

func (obj *Modular) TableName() string {
	return TableName("modular")
}

func GetModulars(pid int) ([]*Modular, error) {
	var modulars []*Modular
	project := Project{Id: pid}
	o := orm.NewOrm()
	_, err := o.QueryTable(TableName("modular")).Filter("Project", project).Filter("IsActive", true).RelatedSel().All(&modulars)
	if err != nil {
		return nil, err
	}
	return modulars, nil
}

func AddModular(obj *Modular) (int64, error) {
	return orm.NewOrm().Insert(obj)
}
