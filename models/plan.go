package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	UserId  int
	User    User
	Courses []CourseGroup `gorm:"many2many:plan_courses;"`
}
