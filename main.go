package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Post("/api/pages", routes.CreatePage)
	app.Get("/api/pages", routes.GetPages)
	app.Get("/api/pages/:id", routes.GetPage)
	app.Put("/api/pages/:id", routes.UpdatePage)
	app.Delete("/api/pages/:id", routes.DeletePage)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	//app.Get("/api", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
