package src

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("Before saving User...")
	if u.Password != "" {
		hash, err := HashPassword(u.Password)
		if err != nil {
		   return nil
		}
		tx.Statement.SetColumn("Password", hash)
	}
	return
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

type School struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
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
	db.AutoMigrate(&User{}, &School{})
	return *db
}
