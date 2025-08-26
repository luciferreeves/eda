package router

import (
	"eda/controllers"

	"github.com/gofiber/fiber/v2"
)

func Initialise(router *fiber.App) {
	router.Static("/static", "./static")
	router.Get("/", controllers.HomeController)
}
