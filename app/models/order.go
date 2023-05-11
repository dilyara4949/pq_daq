package models

import (
	"time"
)

type Order struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	UserID    uint
	ProductId uint
	// User         User
	// Products     Product `gorm:"many2many:order_products;"`
	PaymentStatus string  `gorm:"default:false"`
	TotalPrice    float64
}
