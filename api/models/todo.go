package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       string     `gorm:"not null;" json:"title"`
	Description string     `gorm:"null;" json:"description"`
	DueDate     time.Time  `gorm:"null;" json:"due_date"`
	Priority    string     `gorm:"null;" json:"priority"`
	Status      string     `gorm:"null;" json:"status"`
	User        *User      `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint       `gorm:"not null;" json:"user_id"`
}

func (data *Todo) Prepare() {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Status = "CREATED"
	data.Priority = "HIGH"
}

func (data *Todo) Validate() error {
	if data.Title == "" {
		return errors.New("Title is required")
	}

	if data.UserID == 0 {
		return errors.New("UserID is required")
	}
	return nil
}

func (data *Todo) Save(db *gorm.DB) (*Todo, error) {
	var err error

	err = db.Debug().Model(&Todo{}).Create(&data).Error

	if err != nil {
		return &Todo{}, err
	}
	return data, nil
}

func (data *Todo) Find(db *gorm.DB, id uint) (*Todo, error) {
	var err error
	err = db.Debug().Model(&Todo{}).Preload("User").Where("id=?", id).Take(data).Error

	if err != nil {
		return &Todo{}, err
	}

	return data, nil
}

func (data *Todo) FindAll(db *gorm.DB) (*[]Todo, error) {
	var err error
	datas := []Todo{}
	err = db.Debug().Model(&Todo{}).Preload("User").Order("id desc").Find(&datas).Error

	if err != nil {
		return &[]Todo{}, err
	}
	return &datas, nil
}

func (data *Todo) Update(db *gorm.DB, id int64) (*Todo, error) {
	var err error
	err = db.Debug().Model(&Todo{}).Where("id=?", id).Updates(data).Error
	if err != nil {
		return &Todo{}, err
	}
	return data, nil
}

func (data *Todo) Delete(db *gorm.DB, id uint) (*Todo, error) {
	var err error
	err = db.Debug().Model(&Todo{}).Where("id=?", id).Delete(data).Error
	if err != nil {
		return &Todo{}, err
	}
	return data, nil
}
