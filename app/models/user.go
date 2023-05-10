package models

import (
	"time"
	// "gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)
type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Token     string     `gorm:"-" json:"token,omitempty"`
	Username  string     `gorm:"unique" json:"username"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"password"`
}