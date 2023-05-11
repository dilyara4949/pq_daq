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
	// UserID      uint
	// User        User
	CategoryID  uint
	Category  Category
	Name        string
	Price       float64
	Stock       int
	Description string
	Rating float64
}

type ProductBodyParam struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	CategoryID   string `json:"category_id,omitempty" validate:"required"`
	Price       uint   `json:"price,omitempty"`
}