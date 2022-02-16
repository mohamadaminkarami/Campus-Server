package src

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type User struct {
	gorm.Model
	StudentNumber string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	Email         string
	EntranceYear  int
	Rand          int
}

type School struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

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

func InitDB() gorm.DB {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB_NAME") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=" + os.Getenv("POSTGRES_SSL_MODE") +
		" TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{}, &School{}, &Course{})
	if err != nil {
		panic(err)
	}
	return *db
}
