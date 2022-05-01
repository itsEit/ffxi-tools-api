package routes

import (
	"com/ffxi-tools/controllers"

	"github.com/gofiber/fiber/v2"
)

func ItemRoute(app *fiber.App) {
	app.Post("/item", controllers.CreateItem)
	app.Get("/item/:itemId", controllers.GetAItem)
	app.Put("/item/:itemId", controllers.EditAItem)
	app.Delete("/item/:itemId", controllers.DeleteAItem)
	app.Get("/item", controllers.GetAllItems)
}
