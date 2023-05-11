package models

import (
	// "gorm.io/gorm"
	"time"
)

type Rating struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	// UserID    uint
	// User      User
	ProductID uint
	Product   Product
	Amount    int
	Sum       int
}

type RatingRep struct {
	ProductId   uint
	Rating    float64
}