package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./src/templates", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views"), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("pages/index", fiber.Map{
			"Title":       "HIiiii",
			"CurrentTime": 30,
		}, "layout")
	})

	app.Get("/nested", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/nested/index", fiber.Map{
			"Title":       "Hello, World!",
			"CurrentTime": 30,
		}, "layout")
	})

	app.Post("/components/clock", func(c *fiber.Ctx) error {
		intTime, _ := strconv.Atoi(c.FormValue("clock-time"))
		return c.Render("components/clock", fiber.Map{

			"CurrentTime": intTime + 10,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
