package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucas2500/fiber-rest-api/book"
	"github.com/lucas2500/fiber-rest-api/book/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MainPage struct {
	Result string `json:"result"`
}

func main() {
	app := fiber.New()

	InitDatabase()

	app.Get("/", HelloStranger)
	SetupRoutes(app)

	app.Listen(":3000")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Put("/api/book/:id", book.UpdateBook)
	app.Delete("/api/book/:id", book.DeleteBook)
}

func HelloStranger(c *fiber.Ctx) error {

	var res MainPage

	res.Result = "Hello stranger!! Essa é a página principal da API de livros"

	return c.JSON(res)
}

func InitDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		panic("Houve um erro ao se conectar com o banco de dados")
	}

	fmt.Println("Conexao efetuada com sucesso!!")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Migration executada com sucesso!!")
}
