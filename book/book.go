package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas2500/fiber-rest-api/book/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Titulo    string  `json:"titulo"`
	Autor     string  `json:"autor"`
	Avaliacao float64 `json:"avaliacao"`
}

type Response struct {
	Result string `json:"result"`
}

func GetBooks(c *fiber.Ctx) error {

	db := database.DBConn
	var books []Book
	var res Response

	db.Find(&books)

	fmt.Println(books)

	if len(books) == 0 {
		res.Result = "Não há livros cadastrados!!"
		return c.Status(404).JSON(res)
	}

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn
	var book Book
	var res Response

	db.Find(&book, id)

	if book.Titulo == "" {
		res.Result = "Nenhum livro foi encontrado com o id informado!!"
		return c.Status(404).JSON(res)
	}

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {

	db := database.DBConn

	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err)
	}

	db.Create(&book)

	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {

	id := c.Params("id")
	var res Response
	db := database.DBConn

	book := new(Book)

	db.First(&book, id)

	if book.Titulo == "" {
		res.Result = "Nenhum livro foi encontrado com o id informado!!"
		return c.Status(404).JSON(res)
	}

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err)
	}

	db.Save(&book)

	res.Result = "Livro atualizado com sucesso!!"
	return c.Status(200).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn

	var book Book
	var res Response

	db.First(&book, id)

	if book.Titulo == "" {
		res.Result = "Nenhum livro foi encontrado com o id informado!!"
		return c.Status(404).JSON(res)
	}

	db.Delete(&book)

	res.Result = "Livro deletado com sucesso!!"
	return c.JSON(res)
}
