package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	ID       uint    
	Name     string   
}
