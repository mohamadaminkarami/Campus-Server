package models

import (
	"gorm.io/gorm"
	"time"
)

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

type CourseGroup struct {
	gorm.Model
	ProfessorId int
	Professor   Professor
	CourseId    int
	Course      Course
	GroupNumber int `gorm:"not null"`
	Capacity    int `gorm:"not null"`
	ExamDate    time.Time
	Detail      string
	Schedule    []Schedule
}

type Schedule struct {
	gorm.Model
	Start         float64 `gorm:"not null"`
	End           float64 `gorm:"not null"`
	Day           int     `gorm:"not null"`
	CourseGroupId int
	CourseGroup   CourseGroup
}
