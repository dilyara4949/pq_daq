package db

import ( 
	"fmt"
	"log"
	"os"

	"github.com/dilyara4949/pq_daq/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

var DB Database 

func ConnectDb() {

	dbHost := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s ", dbHost, userName, password, dbName, dbPort, sslMode)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s ", "localhost", "postgres", "12345", "pqdaq", "5432", "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database %v\n, dsn: %s", err, dsn)
	}

	err = db.AutoMigrate(&models.User{}, &models.Order{}, &models.Category{}, &models.Product{}, &models.Comment{}, &models.Rating{})
	// , &models.Order{}, &models.Category{}, &models.Product{}, &models.Comment{}, &models.Rating{}


	if err != nil {
		
		panic(err)
	}

	DB = Database {
		Db: db,
	}
}