package controllers

import (
	"eda/config"
	"eda/utils/github"
	"eda/utils/shortcuts"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HomeController(ctx *fiber.Ctx) error {
	ctx.Locals("Title", "Home")

	contributionGraph, err := github.GetContributionGraph()
	if err != nil {
		log.Printf("failed to get contribution graph: %v\n", err)
	}
	return shortcuts.Render(ctx, config.TEMPLATE_HOME, fiber.Map{
		"User":                 github.GetProfile(),
		"ContributionCalendar": contributionGraph,
	})
}
