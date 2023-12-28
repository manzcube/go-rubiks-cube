package controllers

import (
	"net/http"
	"server/logic/models"
	"server/logic/utils"

	"github.com/gin-gonic/gin"
)

// This function read params and decides which is the right turn to do
func TurnHandler(c *gin.Context) {
	// Take params to know which turn func to execute
	direction := c.Param("direction")

	var turnFunc func(models.Cube) models.Cube

	switch direction {
	case "r":
		turnFunc = turnR
	case "r-prime":
		turnFunc = turnRPrime
	case "l":
		turnFunc = turnL
	case "l-prime":
		turnFunc = turnLPrime
	case "u":
		turnFunc = turnU
	case "u-prime":
		turnFunc = turnUPrime
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid turn direction"})
		return
	}
	
	// Execute cube turn generic func
	HandleCubeTurn(c, turnFunc)
}

// This func handles gin context and request response JSON object
func HandleCubeTurn(c *gin.Context, turnFunc func(models.Cube) models.Cube) {	
	// Define a struct to match the JSON data structure sent from the frontend
	var requestData models.Cube

	// Parse the JSON data from the request body into requestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Take new Cube data
	turnedCube := turnFunc(requestData)

	// Send new data
	c.IndentedJSON(http.StatusOK, turnedCube)
}

// // GET controllers

// Render default cube state
func RenderCube() models.Cube {
	// Create an instance of the Rubik's Cube
	var cube models.Cube

	// Set colors
	for face := 0; face < 6; face++ {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				color := utils.GetColor(face)
				cube[face][row][col] = color
			}
		}
	}
	return cube
}

// // POST turning controllers

// Turn R
func turnR(cube models.Cube) models.Cube {
	// We create a new Cube object instance
	newCube := cube

	// Perfom the turn
	for face := 0; face < 6; face++ {
		if face != 1 && face != 3 {
			for row := 0; row < 3; row++ {
				newCube[face][row][2] = cube[utils.GetFaceForVerticalTurnClockwise(face)][row][2]
			}
		}
	}
	return newCube
}
 
// Turn R Prime
func turnRPrime(cube models.Cube) models.Cube {
	// We create a new Cube object
	newCube := cube

	// Set colors
	for face := 0; face < 6; face++ {
		if face != 1 && face != 3 {
			for row := 0; row < 3; row++ {
				newCube[face][row][2] = cube[utils.GetFaceForVerticalTurnCounterClockwise(face)][row][2]
			}
		}
	}
	return newCube
}

// Turn L
func turnL(cube models.Cube) models.Cube {
	// We create a new Cube object instance
	newCube := cube

	// Perfom the turn
	for face := 0; face < 6; face++ {
		if face != 1 && face != 3 {
			for row := 0; row < 3; row++ {
				newCube[face][row][0] = cube[utils.GetFaceForVerticalTurnCounterClockwise(face)][row][0]
			}
		}
	}
	return newCube
}

// Turn L Prime
func turnLPrime(cube models.Cube) models.Cube {
	// We create a new Cube object
	newCube := cube

	// Set colors
	for face := 0; face < 6; face++ {
		if face != 1 && face != 3 {
			for row := 0; row < 3; row++ {
				newCube[face][row][0] = cube[utils.GetFaceForVerticalTurnClockwise(face)][row][0]
			}
		}
	}
	return newCube
}

// Turn U
func turnU(cube models.Cube) models.Cube {
	// We create a new Cube object instance
	newCube := cube

	// Perfom the turn
	for face := 0; face < 6; face++ {
		if face != 0 && face != 4 {
			for row := 0; row < 3; row++ {
				newCube[face][0][row] = cube[utils.GetFaceForHorizontalTurnCounterClockwise(face)][0][row]
			}
		}
	}
	return newCube
}

// Turn U Prime
func turnUPrime(cube models.Cube) models.Cube {
	// We create a new Cube object
	newCube := cube

	// Set colors
	for face := 0; face < 6; face++ {
		if face != 0 && face != 4 {
			for row := 0; row < 3; row++ {
				newCube[face][0][row] = cube[utils.GetFaceForHorizontalTurnClockwise(face)][0][row]
			}
		}
	}
	return newCube
}
