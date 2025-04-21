package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("/api/v1") // Define the main API group
	{
		info := main.Group("/wisdom/info") // Define the info group
		{
			info.GET("/", controllers.ShowInfo) // Info route
		}

		wisdom := main.Group("/wisdom/dispense") // Define the wisdom group
		{
			wisdom.GET("/", controllers.ShowWisdom) // Wisdom route
		}
	}

	return router // Return the configured router
}
