package routers

import (
	"github.com/gofiber/fiber/v2"
)

func handler_for_link_api(c *fiber.Ctx) error {
	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "LinkPad",
	})
}

func AddRoutersForLink(router fiber.Router) {
	router.Get("/link", handler_for_link_api)
}
