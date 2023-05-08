package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	ID            uint    
	User          User 
	Product       []Product 
	PaymentStatus string  
	TotalPrice    uint     
}
