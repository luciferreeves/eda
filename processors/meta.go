package processors

import "github.com/gofiber/fiber/v2"

const defaultTitle = "default"

func MetaContextProcessor(ctx *fiber.Ctx) error {
	ctx.Locals("Title", defaultTitle)
	return ctx.Next()
}
