package main

import (
	"fmt"
	"log"

	"github.com/akshay-singla/employee-crud-api/config"
	"github.com/akshay-singla/employee-crud-api/db"
	"github.com/akshay-singla/employee-crud-api/employee/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.GetEnvConfig()

	// Initialize database connection
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// set DB Configure
	err = db.Configure(conn, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// verify the database connection
	err = db.Verify(conn)
	if err != nil {
		log.Fatal(err)
	}

	db.CreateTables(conn)

	svc := service.NewService(conn)
	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.POST("/employee", svc.CreateEmployee)
	r.GET("/employee", svc.ListEmployee)
	r.GET("/employee/:id", svc.GetEmployeeByID)
	r.PUT("/employee/:id", svc.UpdateEmployee)
	r.DELETE("/employee/:id", svc.DeleteEmployee)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(r.Run(":8080"))
}
