package routes

import (
	"hot-or-not/backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/personas", handlers.GetPersonas)
		// This is the line that needs to be present and correct:
		api.POST("/personas/:id/toggle-like", handlers.ToggleLike) // <-- Make sure this line is here!
	}
}