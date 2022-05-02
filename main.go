package main

import (
	"com/ffxi-tools/configs"
	"com/ffxi-tools/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)
	routes.ItemRoute(router)

	router.Run("localhost:6000")
}
