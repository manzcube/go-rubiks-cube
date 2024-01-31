package routes

import (
	"net/http"
	"server/logic/controllers"
	"server/logic/models"

	"github.com/gin-gonic/gin"
)

func RegisterCubeRoutes(router *gin.Engine) {
	router.GET("/", func (c *gin.Context) {
		cube := controllers.RenderCubeTensor()
		c.IndentedJSON(http.StatusOK, cube)
	})
	router.POST("/turn/:direction", TurnHandler)
}

// // This function read params and decides which is the right turn to do
func TurnHandler(c *gin.Context) {
	// Take params to know which turn func to execute
	direction := c.Param("direction")

	var turnFunc func(models.Cube) models.Cube

	switch direction {
	case "r":
		turnFunc = controllers.TurnR
	case "r-prime":
		turnFunc = controllers.TurnRPrime
	case "l":
		turnFunc = controllers.TurnL
	case "l-prime":
		turnFunc = controllers.TurnLPrime
	case "u":
		turnFunc = controllers.TurnU
	case "u-prime":
		turnFunc = controllers.TurnUPrime
	case "d":
		turnFunc = controllers.TurnD
	case "d-prime":
		turnFunc = controllers.TurnDPrime
	case "f":
		turnFunc = controllers.TurnF
	case "f-prime":
		turnFunc = controllers.TurnFPrime
	case "b":
		turnFunc = controllers.TurnB
	case "b-prime":
		turnFunc = controllers.TurnBPrime
	case "m":
		turnFunc = controllers.TurnM
	case "m-prime":
		turnFunc = controllers.TurnMPrime
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid turn direction"})
		return
	}
	
	// Execute cube turn generic func
	HandleCubeTurn(c, turnFunc)
}

// This func handles gin context and request/response JSON object
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

