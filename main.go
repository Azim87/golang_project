package main

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.Use(gin.Logger())
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
