package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id      int      `orm:"column(id);auto" description:"id"`
	Name    string   `orm:"column(name);unique" description:"项目名称"`
	Host    string   `orm:"column(host);" description:"项目域名"`
	User    *User    `orm:"null;rel(fk);on_delete(set_null)" description:"关联用户"`
	Modular *Modular `orm:"reverse(one)"`
	BaseModel
}

func (obj *Project) TableName() string {
	return TableName("project")
}

// 添加项目
func AddProject(obj *Project) (int64, error) {
	return orm.NewOrm().Insert(obj)
}

// 根据项目名称查询项目
func GetProjectByName(name string) (*Project, error) {
	project := new(Project)
	o := orm.NewOrm()
	err := o.QueryTable(TableName("project")).Filter("name", name).One(project)
	if err != nil {
		return nil, err
	}
	return project, err
}

// 获取所有项目
func GetAllProject(uid int) ([]Project, error) {
	var ProList []Project
	o := orm.NewOrm()
	_, err := o.QueryTable(TableName("project")).Filter("IsActive", true).Filter("User", uid).RelatedSel().All(&ProList)
	if err != nil {
		return nil, err
	}
	return ProList, nil
}

func DeleteProject(id int) bool {
	project := &Project{Id: id}
	o := orm.NewOrm()
	if _, err := o.Delete(project); err == nil {
		return true
	} else {
		return false
	}
}
