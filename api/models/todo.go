package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"not null;" json:"title"`
	Description string    `gorm:"null;" json:"description"`
	DueDate     time.Time `gorm:"null;" json:"dueDate"`
	Priority    string    `gorm:"null;" json:"priority"`
	Status      string    `gorm:"null;" json:"status"`
	User        *User     `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint      `gorm:"not null;" json:"userId"`
	ShareWith   []uint    `gorm:"null;" json:"shareWith"`
}
