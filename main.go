package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas2500/fiber-rest-api/book"
)

func main() {
	app := fiber.New()

	app.Get("/", HelloStranger)
	SetupRoutes(app)

	app.Listen(":3000")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book/:id", book.DeleteBook)
}

func HelloStranger(c *fiber.Ctx) error {

	return c.SendString("Hello stranger!! Essa é a página principal da API de livros")
}
