package migrations

import (
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/dilyara4949/pq_daq/app/models"
	"github.com/dilyara4949/pq_daq/db"
)



func Migrate() {
	Product := models.Product{}
	Category := models.Category{}
	Order := models.Order{}
	User := models.User{}

	db.DB.Db.AutoMigrate(&Product, &Category, &Order,  &User)
}