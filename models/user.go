package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentNumber   string `gorm:"unique;not null"`
	Password        string `gorm:"not null"`
	Email           string
	EntranceYear    int
	TakeCoursesTime int
	SchoolId        int
	School          School `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("before creating User...")
	log.Println(u.TakeCoursesTime, u.Password, u.StudentNumber, u.Email)
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
