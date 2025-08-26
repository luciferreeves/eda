package controllers

import (
	"eda/config"
	"eda/utils/shortcuts"

	"github.com/gofiber/fiber/v2"
)

func HomeController(ctx *fiber.Ctx) error {
	return shortcuts.Render(ctx, config.TEMPLATE_HOME, nil)
}
