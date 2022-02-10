package archive

import "github.com/beego/beego/v2/client/orm"

type Professor struct {
	Id   int    `orm:"column(id);pk;auto"`
	Name string `valid:"Required"`
	Bio  string
}

type Comment struct {
	Id          int        `orm:"column(id);pk;auto"`
	Professor   *Professor `orm:"rel(fk)"`
	User        *User      `orm:"rel(fk)"`
	IsAnonymous bool
	text        string `valid:"Required"`
}

type CommentVotes struct {
	Id       int      `orm:"column(id);pk;auto"`
	Comment  *Comment `orm:"rel(fk)"`
	User     *User    `orm:"rel(fk)"`
	IsUpVote bool
}

func init() {
	orm.RegisterModel(new(Professor), new(Comment), new(CommentVotes))
}
