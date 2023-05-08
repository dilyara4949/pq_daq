package models

import (
	// "gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Comment   string
}
