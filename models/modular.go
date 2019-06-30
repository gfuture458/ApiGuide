package models

type Modular struct {
	Id        int        `orm:"column(id);auto" description:"id"`
	Name      string     `orm:"column(name);unique" description:"模块名称"`
	Version   string     `orm:"column(version)" description:"版本"`
	Prefix    string     `orm:"column(prefix)" description:"路由前缀"`
	Project   *Project   `orm:"null;rel(one);on_delete(set_null)" description:"关联项目"`
	Interface *Interface `orm:"reverse(one)"`
	BaseModel
}

func (obj *Modular) TableName() string {
	return TableName("modular")
}
