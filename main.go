package main

import (
	"awesomeAPI/db"
	"awesomeAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
