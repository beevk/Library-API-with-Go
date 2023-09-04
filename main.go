package main

import (
	"BOOK-API/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CatchAllErrorLogger(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			errorMessage := fmt.Sprintf("Server Error: %v", err)
			fmt.Println(errorMessage)
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		}
	}()
	c.Next()
}

func main() {
	router := gin.Default()

	router.Use(CatchAllErrorLogger)

	routes.RegisterRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
