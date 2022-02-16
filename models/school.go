package models

import (
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}