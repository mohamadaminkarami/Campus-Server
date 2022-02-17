package models

import (
	"time"

	"gorm.io/gorm"
)

func getCourseId(DB gorm.DB, code int) (int) {
	var course Course
	DB.First(&course, "code = ?", code)
	return int(course.ID)
}

func InsertDummyData(DB gorm.DB) {
	// School
	schools := []School{
		{Name: "مهندسی کامپیوتر"}, // 1
		{Name: "مهندسی صنایع"}, // 2
		{Name: "مهندسی برق"}, // 3
		{Name: "مهندسی عمران"}, // 4
		{Name: "مرکز معارف"}, // 5
		{Name: "مرکز تربیت بدنی"}, //6
		{Name: "ریاضی"}, // 7
		{Name: "فیزیک"}, // 8
		{Name: "کارگاه‌ها"}, // 9
		{Name: "مهندسی و علم مواد‍"}, // 10
	}
	DB.Create(&schools)

	// Users
	users := []User{
		{
			SchoolId:        1,
			Email:           "kambiz@gmail.com",
			EntranceYear:    2001,
			StudentNumber:   "98105000",
			Password:        "1234",
			TakeCoursesTime: 1645080899,
		},
		{
			SchoolId:        1,
			Email:           "dambiz@gmail.com",
			EntranceYear:    2021,
			StudentNumber:   "98106000",
			Password:        "1234",
			TakeCoursesTime: 1645080700,
		},
	}
	DB.Create(&users)

	// Professor
	professors := []Professor{
		{Name: "امیرعلی چکاه"}, // 1
		{Name: "استاد رئیسی"}, // 2
		{Name: "استاد مقدسی"}, // 3
		{Name: "علی شریفی زارچی"}, // 4
		{Name: "علی عیسی زاده"}, // 5
		{Name: "استاد جهانگیر"}, // 6
	}
	DB.Create(&professors)

	// Course
	courses := []Course{
		{Code: 40153, Name: "مبانی برنامه سازی", Credit: 3, SchoolId: 1},
		{Code: 40244, Name: "برنامه سازی پیشرفته", Credit: 3, SchoolId: 1},
		{Code: 24011, Name: "فیزیک ۱", Credit: 3, SchoolId: 8},
		{Code: 22015, Name: "ریاضی عمومی ۱", Credit: 4, SchoolId: 7},
		{Code: 40115, Name: "ساختمان های گسسته", Credit: 3, SchoolId: 1},
	}
	DB.Create(&courses)

	var course Course
	DB.First(&course, "code = ?", 40244)
	var coursePre Course
	DB.First(&coursePre, "code = ?", 40153)
	DB.Model(&course).Association("Prerequisites").Append([]Course{coursePre})

	// CourseGroup
	courseGroups := []CourseGroup{
		{ // 1
			ProfessorId: 1,
			CourseId: getCourseId(DB, 40153),
			GroupNumber: 1,
			Capacity: 40,
			ExamDate: time.Now().Add(24*64*time.Hour),
			Detail: "حضوری برگزار می‌شود.",
		},
		{ // 2
			ProfessorId: 2,
			CourseId: getCourseId(DB, 40153),
			GroupNumber: 2,
			Capacity: 40,
			ExamDate: time.Now().Add(24*64*time.Hour),
			Detail: "حضوری برگزار می‌شود.",
		},
		{ // 3
			ProfessorId: 3,
			CourseId: getCourseId(DB, 22015),
			GroupNumber: 1,
			Capacity: 100,
			ExamDate: time.Now().Add(24*60*time.Hour),
			Detail: "حضوری برگزار می‌شود.",
		},
		{ // 4
			ProfessorId: 3,
			CourseId: getCourseId(DB, 22015),
			GroupNumber: 2,
			Capacity: 100,
			ExamDate: time.Now().Add(24*60*time.Hour),
			Detail: "حضوری برگزار می‌شود.",
		},
		{ // 5
			ProfessorId: 5,
			CourseId: getCourseId(DB, 40244),
			GroupNumber: 1,
			Capacity: 50,
			ExamDate: time.Now().Add(24*62*time.Hour),
			Detail: "مجازی برگزار می‌شود.",
		},
	}
	DB.Create(&courseGroups)

	// Schedule
	schedules := []Schedule{
		{Start: 7.5, End: 9, Day: 0, CourseGroupId: 1},
		{Start: 7.5, End: 9, Day: 2, CourseGroupId: 1},
		{Start: 7.5, End: 9, Day: 0, CourseGroupId: 2},
		{Start: 7.5, End: 9, Day: 2, CourseGroupId: 2},
		{Start: 10.5, End: 12, Day: 1, CourseGroupId: 3},
		{Start: 10.5, End: 12, Day: 3, CourseGroupId: 3},
		{Start: 10.5, End: 12, Day: 1, CourseGroupId: 4},
		{Start: 10.5, End: 12, Day: 3, CourseGroupId: 4},
		{Start: 16.5, End: 18, Day: 1, CourseGroupId: 5},
		{Start: 16.5, End: 18, Day: 3, CourseGroupId: 5},
	}
	DB.Create(&schedules)
	
	// Plan
	plans := []Plan{
		{UserId: 1}, {UserId: 1}, {UserId: 2},
	}
	DB.Create(&plans)

	var plan Plan
	DB.First(&plan, 1)
	var selectedCourseGroups []CourseGroup
	DB.Find(&selectedCourseGroups, []int{1, 3 , 5})
	DB.Model(&plan).Association("Courses").Append(&selectedCourseGroups)
}
