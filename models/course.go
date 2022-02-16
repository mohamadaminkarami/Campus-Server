package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Code          int    `gorm:"unique;not null"`
	Credit        int    `gorm:"not null"`
	Syllabus      string
	SchoolId      int
	School        School
	Prerequisites []Course `gorm:"many2many:course_prerequisites;"`
	Corequisites  []Course `gorm:"many2many:course_corequisites;"`
}
