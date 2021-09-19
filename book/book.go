package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Retorna todos os livros cadastrados!!")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Retorna um único livro!!")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("Cadastra um novo livro!!")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Deleta um livro!!")
}
