package models

type Interface struct {
	Id      int      `orm:"column(id);auto" description:"id"`
	Name    string   `orm:"column(name);unique" description:"接口名称"`
	Path    string   `orm:"column(path);" description:"路径"`
	Method  string   `orm:"column(method)" description:"请求方式"`
	Require bool     `orm:"column(require)" description:"必填"`
	Example string   `orm:"column(example)" description:"示例"`
	Type    string   `orm:"column(type)" description:"类型"`
	Desc    string   `orm:"column(desc)" description:"描述"`
	Modular *Modular `orm:"null;rel(one);on_delete(set_null)" description:"关联模块"`
	BaseModel
}

func (obj *Interface) TableName() string {
	return TableName("interface")
}
