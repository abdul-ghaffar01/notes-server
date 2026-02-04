package server

import (
	"abdul-ghaffar01/notes-server/internal/health"
	"abdul-ghaffar01/notes-server/internal/note"

	"github.com/gin-gonic/gin"
)


func HandleRoutes(router *gin.Engine){
	
	// Creating the service
	service := note.NewService()

	// Creating a new handler
	handler := note.NewHandler(service)

	router.GET("/health", health.HealthCheckHandler)

	router.POST("/create", handler.Create)
	
}