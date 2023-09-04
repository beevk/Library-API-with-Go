package main

import (
	"BOOK-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register routes defined in the routes.go file
	routes.RegisterRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
