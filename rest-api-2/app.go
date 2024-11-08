package main

import (
	"example.com/rest-api-2/db"
	"example.com/rest-api-2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run("localhost:8080")
}
