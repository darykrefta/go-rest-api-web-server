package main

import (
	"rest-api/database"
	"rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	routes.RegisterRouters(server)

	server.Run(":8080") // starts the server with localhost  
}

