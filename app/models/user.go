package models

import (
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint  
	Username  string
	FirstName string 
	LastName  string 
	Email     string 
	Password  string 
}
