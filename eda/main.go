package main

import (
	"eda/config"
	"eda/processors"
	"eda/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./templates", ".django")
	engine.Reload(config.Server.IsDevMode)

	server := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: serverErrorHandler,
	})
	server.Use(recover.New())
	server.Use(logger.New())
	server.Use(helmet.New(helmet.Config{
		CrossOriginEmbedderPolicy: "unsafe-none",
	}))
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	processors.Initialise(server)
	router.Initialise(server)

	log.Fatalf("Server failed to start: %v", server.Listen(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}

func serverErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	} else if err != nil {
		msg = err.Error()
	}
	return ctx.Status(code).SendString(msg)
}
