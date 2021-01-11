package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
}
