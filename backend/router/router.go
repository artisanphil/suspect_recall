package router

import (
	"suspectRecall/handlers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(r *gin.Engine) {
	r.GET("/api/person/attributes", handlers.GetItems)
	r.GET("/api/person/image/:filename", handlers.GetImage)
}
