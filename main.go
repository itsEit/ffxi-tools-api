package main

import (
	"com/ffxi-tools/configs"
	"com/ffxi-tools/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)
	routes.ItemRoute(app)

	app.Listen(":6000")
}
