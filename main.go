package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var port = flag.Int("p", 8080, "서버가 Listen할 port 번호를 입력해주세요.")

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	addr := fmt.Sprintf(":%d", *port)
	// use the fiber view engine for rendering engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

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
