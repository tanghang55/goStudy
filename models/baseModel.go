package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var selectDb string ="db"

//db配置初始化
func init() {
	dbHost := beego.AppConfig.String(selectDb + "::db.host")
	dbUser := beego.AppConfig.String(selectDb + "::db.user")
	dbPwd := beego.AppConfig.String(selectDb + "::db.pwd")
	dbName := beego.AppConfig.String(selectDb + "::db.name")
	dbPort := beego.AppConfig.String(selectDb + "::db.port")
	dbTimeZone := beego.AppConfig.String(selectDb + "::db.timezone")
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if dbTimeZone != "" {
		dsn += "&loc" + url.QueryEscape(dbTimeZone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Admin),new(Auth),new(RoleAuth))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TablesName(name string) string {
	return beego.AppConfig.String(selectDb+"::db.prefix") + name
}
