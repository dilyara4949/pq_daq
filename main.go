package main

import (
	"fmt"
	"log"
	"net/http"

	// "net/http"

	"github.com/dilyara4949/pq_daq/db"
	"github.com/dilyara4949/pq_daq/app"
	"github.com/dilyara4949/pq_daq/app/router"
	// repo "github.com/dilyara4949/pq_daq/app/repository"
	// "github.com/gin-gonic/gin"
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
	container := app.Build()
	engine := router.GinInit(container)

	server := &http.Server{
		Addr:    ":8000",
		Handler: engine,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
	

	// userRoutes := r.Group("api/user")
	// {
	// 	userRoutes.GET("/", getUsers)
	// }

}

// func getUsers(c *gin.Context){
// 	s := "lnrlbf"

// 	c.IndentedJSON(http.StatusOK, &s)
// }