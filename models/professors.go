package models

import "github.com/beego/beego/v2/client/orm"

type Professor struct {
	Id   int    `orm:"auto"`
	Name string `valid:"Required"`
	Bio  string
}

type Comment struct {
	Id          int        `orm:"auto"`
	Professor   *Professor `orm:"rel(fk)"`
	User        *User      `orm:"rel(fk)"`
	IsAnonymous bool
	text        string `valid:"Required"`
}

type CommentVotes struct {
	Id       int      `orm:"auto"`
	Comment  *Comment `orm:"rel(fk)"`
	User     *User    `orm:"rel(fk)"`
	IsUpVote bool
}

func init() {
	orm.RegisterModel(new(Professor), new(Comment), new(CommentVotes))
}
