package main

import (
	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/db"
	"rishisingh.in/event-manager/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
