package server

import (
	"abdul-ghaffar01/notes-server/internal/health"

	"github.com/gin-gonic/gin"
)


func HandleRoutes(router *gin.Engine){
	
	router.GET("/health", health.HealthCheckHandler)
}