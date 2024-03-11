package main

import (
	"log"
	"server/logic/constants"
	"server/logic/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize GIN router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Options
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{constants.PublicEndpoint}) // Client endpoint

	// Register Routes
	routes.RegisterCubeRoutes(router)

	// Communicate
	router.Run(":8080")
	log.Printf("GIN server running on port 8080\n")
}
