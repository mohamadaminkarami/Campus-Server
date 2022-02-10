package archive

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id            int    `orm:"column(id);pk;auto"`
	StudentNumber string `valid:"Numeric"`
	Password      string `form:"Password" valid:"Required;MinSize(6)"`
	Email         string `form:"Email" valid:"Required;Email"`
	EntranceYear  int
	Rand          int
	School        *School         `orm:"rel(fk)"`
	PassedCourse  []*PassedCourse `orm:"rel(m2m);rel_through(models.PassedCourse)"`
}

type PassedCourse struct {
	Id       int          `orm:"column(id);pk;auto"`
	User     *User        `orm:"rel(fk)"`
	Course   *CourseGroup `orm:"rel(fk)"`
	isPassed bool
}

func init() {
	orm.RegisterModel(new(User), new(PassedCourse))
}
