package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Name string `gorm:"not null"`
}
