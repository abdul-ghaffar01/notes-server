package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run Creates a server and run it.
func Run() {
	// Port to run the server
	Port := 3000

	// Creating a gin router
	router := gin.Default()

	// Listerning to all the routes
	HandleRoutes();

	// Running the server 
	router.Run(fmt.Sprintf(":%d", Port))
}
