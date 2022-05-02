package routes

import (
	"com/ffxi-tools/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoute(router *gin.Engine) {
	router.GET("/item/:itemId", controllers.GetAItem())
	router.PUT("/item/:itemId", controllers.EditAItem())
	router.DELETE("/item/:itemId", controllers.DeleteAItem())
	router.GET("/item", controllers.GetAllItems())
}
