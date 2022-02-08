package models

import "github.com/beego/beego/v2/client/orm"

type School struct {
	Id   int    `orm:"auto"`
	Name string `valid:"Required"`
}

func init() {
	orm.RegisterModel(new(School))
}
