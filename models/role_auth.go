package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"bytes"
	"strconv"
)

type RoleAuth struct {
	AuthId int `orm:"pk"`
	RoleId int64
}

func init() {
	selectDb = "db" //选库
}
func (r *RoleAuth) TableName() string {
	// db table name
	return TablesName("uc_role_auth")
}

//获取多个
func (r *RoleAuth) RoleAuthGetByIds(roleIds string) (authIds string, err error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(r.TableName())
	ids := strings.Split(roleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	authIds = strings.TrimRight(b.String(), ",")
	return authIds, nil
}
