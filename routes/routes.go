package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func registerDefaultRoutes(router *gin.Engine) {
	router.HEAD("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
	router.GET("/health", healthCheck)
}

// RegisterRoutes Exported function to register routes
func RegisterRoutes(router *gin.Engine) {
	registerBookRoutes(router)
	registerDefaultRoutes(router)
}
