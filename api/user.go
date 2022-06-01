package api

import (
	"github.com/gofiber/fiber/v2"
)

func handler_for_user_api(c *fiber.Ctx) error {
	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "LinkPad",
	})
}

func AddRoutersForUser(app *fiber.App) {
	app.Get("/user", handler_for_user_api)
}
