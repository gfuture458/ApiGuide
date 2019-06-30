package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func Init() {
	db_host := beego.AppConfig.String("db.host")
	db_port := beego.AppConfig.String("db.port")
	db_user := beego.AppConfig.String("db.user")
	db_name := beego.AppConfig.String("db.name")
	db_pwd := beego.AppConfig.String("db.password")
	timezone := beego.AppConfig.String("db.timezone")
	if db_port == "" {
		db_port = "3306"
	}
	dsn := db_user + ":" + db_pwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Project), new(Modular), new(Interface))
	orm.RunSyncdb("default", false, true)
}
