package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	UUID      string         `json:"uuid" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Gender    string         `json:"gender"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (Contact) TableName() string {
	return "contacts"
}
