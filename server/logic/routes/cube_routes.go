package routes

import (
	"net/http"
	"server/logic/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCubeRoutes(router *gin.Engine) {
	router.GET("/", func (c *gin.Context) {
		cube := controllers.RenderCube()
		c.IndentedJSON(http.StatusOK, cube)
	})
	router.GET("/combinations", func (c *gin.Context) {
		cube := controllers.RenderCubeTensor()
		c.IndentedJSON(http.StatusOK, cube)
	})
	router.POST("/turn/:direction", controllers.TurnHandler)
}


