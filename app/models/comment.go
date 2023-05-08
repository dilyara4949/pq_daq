package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	ID      uint    
	User    string
	Product []string 
	comment string
}
