package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Course struct {
	Id            int    `orm:"auto"`
	Name          string `valid:"Required"`
	Code          int    `valid:"Required"`
	Credit        int    `valid:"Required"`
	Syllabus      string
	School        *School   `orm:"rel(fk)"`
	Prerequisites []*Course `orm:"rel(m2m);rel_through(models.CourseRequisites)"`
	Corequisites  []*Course `orm:"rel(m2m);rel_through(models.CourseRequisites)"`
}

type CourseGroup struct {
	Id          int        `orm:"auto"`
	Professor   *Professor `orm:"rel(fk)"`
	Course      *Course    `orm:"rel(fk)"`
	GroupNumber int        `valid:"Required"`
	Year        int        `valid:"Required"`
	Term        int        `valid:"Required"`
	Capacity    int        `valid:"Required"`
	ExamDate    orm.DateTimeField
	Schedule    []*Schedule `orm:"reverse(many)"`
	Detail      string
}

type CourseRequisites struct {
	Id         int     `orm:"auto"`
	Course     *Course `orm:"rel(fk)"`
	Requisites *Course `orm:"rel(fk)"`
}

type Schedule struct {
	Id          int           `orm:"auto"`
	Start       orm.TimeField `valid:"Required"`
	End         orm.TimeField `valid:"Required"`
	Day         string        `valid:"Required"`
	CourseGroup *CourseGroup  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Course), new(CourseGroup), new(CourseRequisites), new(Schedule))
}
