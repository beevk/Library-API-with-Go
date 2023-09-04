package routes

import (
	"BOOK-API/handlers"
	"github.com/gin-gonic/gin"
)

func registerBookRoutes(router *gin.Engine) {
	api := router.Group("/books")

	api.GET("/", handlers.GetBooks)
	api.GET("/:id", handlers.GetSingleBook)
	api.DELETE("/:id", handlers.DeleteBook)
	api.POST("/", handlers.CreateBook)
	api.PATCH("/rent", handlers.RentABook)
	api.PATCH("/return", handlers.ReturnABook)
}
