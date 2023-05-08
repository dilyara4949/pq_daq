package main

import (
	"fmt"
	"log"
	"net/http"

	// "net/http"

	"github.com/dilyara4949/pq_daq/db"
	// repo "github.com/dilyara4949/pq_daq/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Main Application Starts")
	loadenv()
	db.ConnectDb()
	r := gin.Default()

	

	userRoutes := r.Group("api/user")
	{
		userRoutes.GET("/", getUsers)
	}

	r.Run(":8000")
}

func getUsers(c *gin.Context){
	s := "lnrlbf"

	c.IndentedJSON(http.StatusOK, &s)
}