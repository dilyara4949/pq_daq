package models

import (
	"time"

	// "gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	UserID      uint
	User        User
	Categories  []*Category `gorm:"many2many:product_categories;"`
	Name        string
	Price       float64
	Stock       int
	Description string
}