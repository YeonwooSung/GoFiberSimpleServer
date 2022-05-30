package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

var port = flag.Int("p", 8080, "서버가 Listen할 port 번호를 입력해주세요.")

func HttpServerErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	addr := fmt.Sprintf(":%d", *port)
	// use the fiber view engine for rendering engine
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: HttpServerErrorHandler,
	})
	// use recover for error handling
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		//TODO
		return c.Render("main", fiber.Map{
			"Title": "LinkPad",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})
	log.Printf("Server is listening %d", *port)
	log.Fatal(app.Listen(addr))
}
