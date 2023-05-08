package models

import (
	"github.com/jinzhu/gorm"
)

type Rating struct {
	gorm.Model
	ID      uint  
	User    string
	Product string
	Amount  uint
	Sum     uint
}
