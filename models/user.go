package models

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentNumber string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	Email         string
	EntranceYear  int
	Rand          int
	SchoolId      int `gorm:"foreignKey:SchoolRefer;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
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
