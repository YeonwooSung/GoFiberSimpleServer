package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
)

func DefineRedirectRules(app *fiber.App) {
	// app.Get("/user", handler_for_user_api)
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/redirects": "/",
			"/linkpad":   "/",
		},
		StatusCode: 301,
	}))
}

func handler_for_redirect_api(c *fiber.Ctx) error {
	//TODO cookie parsing, session storing, etc
	return c.Redirect("/")
}

func AddRoutersForRedirect(app *fiber.App) {
	app.Get("/redirect", handler_for_redirect_api)
}
