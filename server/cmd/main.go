package main

import (
	"fmt"

	"server/logic/models"
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
	router.SetTrustedProxies([]string{"127.0.0.1:3000"})

	// Register Routes
	routes.RegisterCubeRoutes(router)

	// Communicate
	fmt.Printf("GIN server running on port 8080\n")
	router.Run("localhost:8080")
}


// Print cube func 
func printCube(cube models.Cube) models.Cube {
	for i := 0; i < len(cube); i++ {
		fmt.Print("\n")
		for j := 0; j < len(cube[i]); j++ {
			fmt.Println(cube[i][j])
		}
	}

	return cube
}
