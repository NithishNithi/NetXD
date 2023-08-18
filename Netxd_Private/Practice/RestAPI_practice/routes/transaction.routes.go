// Define all the routes.
package routes

import (
	"Rest-API/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine, controller controllers.ProfileController){
	router.POST("/api/profile/create", controller.CreateTransaction)
	// router.GET("/api/profile/:id", controller.GetTransaction)
}