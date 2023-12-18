package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cube being an array of 6 3x3 matrices
type EmptyCube [6][3][3]int
type Cube [6][3][3]string

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

	// Routes
	router.GET("/", func(c *gin.Context) {
		cube := renderCube()
		c.IndentedJSON(http.StatusOK, cube)
	})

	router.POST("turn-r", func(c *gin.Context) {
		cube := renderCube()
		turnedCube := turnR(cube)
		c.IndentedJSON(http.StatusOK, turnedCube)
	})

	// Communicate
	fmt.Printf("GIN server running on port 8080\n")
	router.Run("localhost:8080")
}

func renderCube() Cube {
	// Create an instance of the Rubik's Cube
	var cube Cube

	// Set colors
	for face := 0; face < 6; face++ {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				color := getColor(face)
				cube[face][row][col] = color
			}
		}
	}

	return cube
}

// Get Color by piece code num
func getColor(code int) string {
	colorMap := map[int]string {
		0: "W",
		1: "O",
		2: "G",
		3: "R",
		4: "Y",
		5: "B",
	}
	return colorMap[code]
}

// Turn R
func turnR(cube Cube) Cube {
	// We create a new Cube object
	newCube := renderCube()

	// Set colors
	for face := 0; face < 6; face++ {
		if face != 1 && face != 3 {
			for row := 0; row < 3; row++ {
				newCube[face][row][2] = cube[getFaceForTurnR(face)][row][2]
			}
		}
	}
	return newCube
}

// get correct face for turn R
func getFaceForTurnR(face int) int {
	switch face {
    case 0:
        return 2
	case 2:
        return 4
	case 4:
        return 5
	case 5:
        return 0
	default:
		return face
    }
}
	


// Print cube func 
// func printCube(cube Cube) Cube {
// 	for i := 0; i < len(cube); i++ {
// 		fmt.Print("\n")
// 		for j := 0; j < len(cube[i]); j++ {
// 			fmt.Println(cube[i][j])
// 		}
// 	}

// 	return cube
// }



// Get Color by piece code num
// func getColorCode(code string) int {
// 	colorMap := map[string]int {
// 		"G": 0,
// 		"W": 1,
// 		"B": 2,
// 		"Y": 3,
// 		"O": 4,
// 		"R": 5,
// 	}
// 	return colorMap[code]
// }


