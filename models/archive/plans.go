package archive

import "github.com/beego/beego/v2/client/orm"

type Plan struct {
	Id      int            `orm:"column(id);pk;auto"`
	User    *User          `orm:"rel(fk)"`
	Courses []*CourseGroup `orm:"rel(m2m);rel_through(models.PassedCourse)"`
	Credits int
}

type PlanCourseGroup struct {
	Id          int          `orm:"column(id);pk;auto"`
	Plan        *Plan        `orm:"rel(fk)"`
	CourseGroup *CourseGroup `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Plan), new(PlanCourseGroup))
}
