package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int
	LoginName  string `form:"username"`
	RealName   string
	Password   string `form:"password"`
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func init() {
	selectDb = "db1"   //选库
}
//按用户名查询信息
func (a *Admin) GetUser(name string) (*Admin, error) {
	err := orm.NewOrm().QueryTable(TablesName("uc_admin")).Filter("login_name", name).One(a)
	if err != nil {
		return nil, err
	}
	return a, err
}

//更新用户信息
func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
