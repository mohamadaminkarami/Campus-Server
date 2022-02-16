package models

import (
	"backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() gorm.DB {
	dsn := "host=" + config.Get("POSTGRES_HOST") +
		" user=" + config.Get("POSTGRES_USER") +
		" password=" + config.Get("POSTGRES_PASSWORD") +
		" dbname=" + config.Get("POSTGRES_DB_NAME") +
		" port=" + config.Get("POSTGRES_PORT") +
		" sslmode=" + config.Get("POSTGRES_SSL_MODE") +
		" TimeZone=" + config.Get("POSTGRES_TIMEZONE")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &School{}, &Course{}, &Professor{}, &CourseGroup{}, &Schedule{})
	return *db
}
