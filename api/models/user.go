package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     string     `gorm:"not null;unique;" json:"email"`
	Password  string     `gorm:"not null;" json:"password"`
}

func (data *User) Prepare() {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
}

func (data *User) Validate() error {
	if data.Email == "" {
		return errors.New("Email is required")
	}

	if data.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func (data *User) Save(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Create(&data).Error
	if err != nil {
		return &User{}, err
	}
	return data, nil
}

func (data *User) FindByEmail(db *gorm.DB, email string) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("email=?", email).Take(data).Error
	if err != nil {
		return &User{}, err
	}
	return data, nil
}

func (data *User) FindById(db *gorm.DB, id int64) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id=?", id).Take(data).Error
	if err != nil {
		return &User{}, err
	}
	return data, nil
}
