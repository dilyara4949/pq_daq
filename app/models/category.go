package models

import (
	"time"

	// "gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Name string `gorm:"unique"`
}