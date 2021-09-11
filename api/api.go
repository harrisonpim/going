package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harrisonpim/going/api/database"
)

func CreateAuthor(c *fiber.Ctx) error {
	db := database.DBConn
	author := new(database.Author)
	if err := c.BodyParser(author); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process author"})
	}
	db.Create(&author)
	return c.Status(201).JSON(author)
}

func GetAllAuthors(c *fiber.Ctx) error {
	db := database.DBConn
	authors := []database.Author{}
	db.Find(&authors)
	return c.JSON(authors)
}

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	author := database.Author{}
	db.Find(&author, id)

	if author.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	}
	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	author := database.Author{}
	update := new(database.Author)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process author"})
	}
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	}
	db.Model(&author).Update(&update)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	author := database.Author{}
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	}
	db.Delete(&author)
	return c.JSON(fiber.Map{"status": 200, "message": "Author successfully deleted"})
}

func CreateArticle(c *fiber.Ctx) error {
	db := database.DBConn
	article := new(database.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process article"})
	}
	db.Create(&article)
	return c.Status(201).JSON(article)
}

func GetAllArticles(c *fiber.Ctx) error {
	db := database.DBConn
	articles := []database.Article{}
	db.Find(&articles)
	return c.JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	article := database.Article{}

	db.Find(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	return c.JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	article := database.Article{}

	update := new(database.Article)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process article"})
	}
	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	db.Model(&article).Update(&update)
	return c.JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	article := database.Article{}

	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	db.Delete(&article)
	return c.JSON(fiber.Map{"status": 200, "message": "Article successfully deleted"})
}

func main() {
	database.Init()

	app := fiber.New()
	app.Use(logger.New())

	app.Post("/articles", CreateArticle)
	app.Get("/articles", GetAllArticles)
	app.Get("/articles/:id", GetArticle)
	app.Put("/articles/:id", UpdateArticle)
	app.Delete("/articles/:id", DeleteArticle)

	app.Post("/authors", CreateAuthor)
	app.Get("/authors", GetAllAuthors)
	app.Get("/authors/:id", GetAuthor)
	app.Put("/authors/:id", UpdateAuthor)
	app.Delete("/authors/:id", DeleteAuthor)

	app.Listen(":3000")
}
