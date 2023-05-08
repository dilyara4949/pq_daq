package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ID          uint     
	User        string   

	Categories  []string 
	
	Name        string   
	Price       uint   
	Stock       int
	Description string
}
