package main

import (
	"BOOK-API/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_CONNECTION_URL")
	fmt.Println("Test:::", dbUrl)
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
