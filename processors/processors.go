package processors

import "github.com/gofiber/fiber/v2"

func Initialise(app *fiber.App) {
	app.Use(MetaContextProcessor)
}
