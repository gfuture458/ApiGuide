package models

import (
	"Api/models"
	"github.com/astaxie/beego/orm"
)

type Interface struct {
	Id      int      `orm:"column(id);auto" description:"id" json:"id"`
	Name    string   `orm:"column(name);unique" description:"接口名称" json:"name"`
	Path    string   `orm:"column(path);" description:"路径" json:"path"`
	Version string   `orm:"column(version)" description:"版本" json:"version"`
	Method  string   `orm:"column(method)" description:"请求方式" json:"method"`
	Require bool     `orm:"column(require)" description:"必填" json:"require"`
	Example string   `orm:"column(example)" description:"示例" json:"example"`
	Type    string   `orm:"column(type)" description:"类型" json:"type"`
	Desc    string   `orm:"column(desc)" description:"描述" json:"desc"`
	Modular *Modular `orm:"null;rel(fk);on_delete(set_null)" description:"关联模块" json:"modular"`
	BaseModel
}

func (obj *Interface) TableName() string {
	return TableName("interface")
}

func AddInterface(inter *Interface) (int64, error) {
	return orm.NewOrm().Insert(inter)
}

func GetInterface(mid int) ([]*Interface, error) {
	var interfaces []*Interface
	modular := models.Modular{Id: mid}
	_, err := orm.NewOrm().QueryTable(TableName("interface")).Filter("Modular", modular).All(&interfaces)
	if err != nil {
		return nil, err
	}
	return interfaces, err
}
