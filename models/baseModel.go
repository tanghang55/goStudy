package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
)

var defaultDB string = "db"
//db配置初始化
func init() {
	dbHost := beego.AppConfig.String(getDb(defaultDB) + ":db.host")
	dbUser := beego.AppConfig.String(getDb(defaultDB) + ":db.user")
	dbPwd := beego.AppConfig.String(getDb(defaultDB) + ":db.pwd")
	dbName := beego.AppConfig.String(getDb(defaultDB) + ":db.name")
	dbPort := beego.AppConfig.String(getDb(defaultDB) + ":db.port")
	dbTimeZone := beego.AppConfig.String(getDb(defaultDB) + ":db.timezone")
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if dbTimeZone != "" {
		dsn += "&loc" + url.QueryEscape(dbTimeZone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Admin))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func getDb(name string) string {
	return name
}

func TablesName(name string) string {
	return beego.AppConfig.String(getDb(defaultDB)+"::db.prefix") + name
}
