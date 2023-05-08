package models

import (
	"time"

	// "gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}