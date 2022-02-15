package src

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentNumber string `gorm:"unique;not null"` 
	Password      string `gorm:"not null"` 
	Email         string
	EntranceYear  int
	Rand          int
}

func InitDB() gorm.DB {
	dsn := "host="+ os.Getenv("POSTGRES_HOST") + 
		   " user="+ os.Getenv("POSTGRES_USER") + 
		   " password="+ os.Getenv("POSTGRES_PASSWORD") + 
		   " dbname=" + os.Getenv("POSTGRES_DB_NAME") +
		   " port=" + os.Getenv("POSTGRES_PORT") +
		   " sslmode="+ os.Getenv("POSTGRES_SSL_MODE") +
		   " TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	return *db
}