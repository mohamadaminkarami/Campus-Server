package models

import (
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func getCourseGroupId(DB gorm.DB, courseId int, groupNumber int) int {
	var courseGroup CourseGroup
	DB.Where(&CourseGroup{CourseId: courseId, GroupNumber: groupNumber}).First(&courseGroup)
	return int(courseGroup.ID)
}

func getProfessorId(DB gorm.DB, name string) int {
	var professor Professor
	DB.First(&professor, "name = ?", name)
	return int(professor.ID)
}

func getSchoolId(DB gorm.DB, name string) int {
	var school School
	DB.First(&school, "name = ?", name)
	return int(school.ID)
}

func getCourseId(DB gorm.DB, code int) int {
	var course Course
	DB.First(&course, "code = ?", code)
	return int(course.ID)
}

func InsertDummyData(users_count int) {
	DB := *GetDB()
	// School
	schools := []School{
		{Name: "مهندسی کامپیوتر"},    // 1
		{Name: "مهندسی صنایع"},       // 2
		{Name: "مهندسی برق"},         // 3
		{Name: "مهندسی عمران"},       // 4
		{Name: "مرکز معارف"},         // 5
		{Name: "مرکز تربیت بدنی"},    //6
		{Name: "ریاضی"},              // 7
		{Name: "فیزیک"},              // 8
		{Name: "کارگاه‌ها"},          // 9
		{Name: "مهندسی و علم مواد‍"}, // 10
		{Name: "مرکز زبان‌ها"},       // 11
	}
	DB.Create(&schools)

	// Users
	users := []User{}
	times := []int{1655528649, 1655538649, 1655548649, 1655558649}
	for i := 0; i <= users_count; i++ {
		rand.Seed(time.Now().UnixNano())
		stid := 98101200 + i
		user := User{
			SchoolId:        getSchoolId(DB, "مهندسی کامپیوتر"),
			Email:           "kambiz@gmail.com",
			EntranceYear:    rand.Intn(10) + 1390,
			StudentNumber:   strconv.Itoa(stid),
			Password:        "1234",
			TakeCoursesTime: times[rand.Intn(len(times))],
		}
		users = append(users, user)
	}

	DB.Create(&users)

	// Professor
	professors := []Professor{
		{Name: "امیرعلی چکاه"},     // 1
		{Name: "استاد رئیسی"},      // 2
		{Name: "استاد مقدسی"},      // 3
		{Name: "علی شریفی زارچی"},  // 4
		{Name: "مرتضی عیسی زاده"},  // 5
		{Name: "استاد جهانگیر"},    // 6
		{Name: "استاد جواهریان"},   // 7
		{Name: "استاد ضرابی زاده"}, // 8
	}
	DB.Create(&professors)

	// Course
	courses := []Course{
		{Code: 40153, Name: "مبانی برنامه سازی", Credit: 3, SchoolId: getSchoolId(DB, "مهندسی کامپیوتر")},
		{Code: 40244, Name: "برنامه سازی پیشرفته", Credit: 3, SchoolId: getSchoolId(DB, "مهندسی کامپیوتر")},
		{Code: 24011, Name: "فیزیک ۱", Credit: 3, SchoolId: getSchoolId(DB, "فیزیک")},
		{Code: 22015, Name: "ریاضی عمومی ۱", Credit: 4, SchoolId: getSchoolId(DB, "ریاضی")},
		{Code: 33018, Name: "کارگاه عمومی", Credit: 1, SchoolId: getSchoolId(DB, "کارگاه‌ها")},
		{Code: 40108, Name: "کارگاه کامپیوتر", Credit: 1, SchoolId: getSchoolId(DB, "کارگاه‌ها")},
		{Code: 31119, Name: "آشنایی با ادبیات فارسی", Credit: 3, SchoolId: getSchoolId(DB, "مرکز زبان‌ها")},
		{Code: 40115, Name: "ساختمان های گسسته", Credit: 3, SchoolId: getSchoolId(DB, "مهندسی کامپیوتر")},
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
			ProfessorId: getProfessorId(DB, "امیرعلی چکاه"),
			CourseId:    getCourseId(DB, 40153),
			GroupNumber: 1,
			Capacity:    40,
			ExamDate:    time.Now().Add(24 * 64 * time.Hour),
			Detail:      "حضوری برگزار می‌شود.",
		},
		{ // 2
			ProfessorId: getProfessorId(DB, "استاد رئیسی"),
			CourseId:    getCourseId(DB, 40153),
			GroupNumber: 2,
			Capacity:    40,
			ExamDate:    time.Now().Add(24 * 64 * time.Hour),
			Detail:      "حضوری برگزار می‌شود.",
		},
		{ // 3
			ProfessorId: getProfessorId(DB, "استاد مقدسی"),
			CourseId:    getCourseId(DB, 22015),
			GroupNumber: 1,
			Capacity:    100,
			ExamDate:    time.Now().Add(24 * 60 * time.Hour),
			Detail:      "حضوری برگزار می‌شود.",
		},
		{ // 4
			ProfessorId: getProfessorId(DB, "استاد مقدسی"),
			CourseId:    getCourseId(DB, 22015),
			GroupNumber: 2,
			Capacity:    100,
			ExamDate:    time.Now().Add(24 * 60 * time.Hour),
			Detail:      "حضوری برگزار می‌شود.",
		},
		{ // 5
			ProfessorId: getProfessorId(DB, "مرتضی عیسی زاده"),
			CourseId:    getCourseId(DB, 40244),
			GroupNumber: 1,
			Capacity:    50,
			ExamDate:    time.Now().Add(24 * 62 * time.Hour),
			Detail:      "مجازی برگزار می‌شود.",
		},
		{ // 6
			ProfessorId: getProfessorId(DB, "استاد ضرابی زاده"),
			CourseId:    getCourseId(DB, 40108),
			GroupNumber: 1,
			Capacity:    20,
			ExamDate:    time.Now().Add(24 * 58 * time.Hour),
			Detail:      "مجازی برگزار می‌شود.",
		},
		{ // 7
			ProfessorId: getProfessorId(DB, "استاد ضرابی زاده"),
			CourseId:    getCourseId(DB, 40108),
			GroupNumber: 2,
			Capacity:    25,
			ExamDate:    time.Now().Add(24 * 58 * time.Hour),
			Detail:      "مجازی برگزار می‌شود.",
		},
		{ // 8
			ProfessorId: getProfessorId(DB, "استاد جواهریان"),
			CourseId:    getCourseId(DB, 31119),
			GroupNumber: 1,
			Capacity:    35,
			ExamDate:    time.Now().Add(24 * 69 * time.Hour),
			Detail:      "مجازی برگزار می‌شود.",
		},
	}
	DB.Create(&courseGroups)

	// Schedule
	schedules := []Schedule{
		{Start: 7.5, End: 9, Day: 0,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40153), 1)},
		{Start: 7.5, End: 9, Day: 2,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40153), 1)},
		{Start: 7.5, End: 9, Day: 0,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40153), 2)},
		{Start: 7.5, End: 9, Day: 2,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40153), 2)},
		{Start: 10, End: 12, Day: 1,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 22015), 1)},
		{Start: 10, End: 12, Day: 3,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 22015), 1)},
		{Start: 10, End: 12, Day: 0,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 22015), 2)},
		{Start: 10, End: 12, Day: 2,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 22015), 2)},
		{Start: 16.5, End: 18, Day: 1,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40244), 1)},
		{Start: 16.5, End: 18, Day: 3,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40244), 1)},
		{Start: 9, End: 12, Day: 4,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40108), 1)},
		{Start: 13, End: 16, Day: 4,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 40108), 2)},
		{Start: 15, End: 16.30, Day: 0,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 31119), 1)},
		{Start: 15, End: 16.30, Day: 2,
			CourseGroupId: getCourseGroupId(DB, getCourseId(DB, 31119), 1)},
	}
	DB.Create(&schedules)

	// One plan is created for each user with AfterCreate
	var dbUsers []User
	DB.Find(&dbUsers)
	for _, user := range dbUsers {
		var plan Plan
		DB.Where("user_id = ?", user.ID).First(&plan)
		var selectedCourseGroups []CourseGroup
		rand.Seed(time.Now().UnixNano())
		DB.Find(&selectedCourseGroups,
			[]int{
				getCourseGroupId(DB, getCourseId(DB, 40153), rand.Intn(2)),
				getCourseGroupId(DB, getCourseId(DB, 22015), rand.Intn(2)),
				getCourseGroupId(DB, getCourseId(DB, 40244), 1),
				getCourseGroupId(DB, getCourseId(DB, 40108), 2),
				getCourseGroupId(DB, getCourseId(DB, 31119), 1),
			})
		DB.Model(&plan).Association("Courses").Append(&selectedCourseGroups)
	}
}
