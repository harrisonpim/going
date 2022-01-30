package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harrisonpim/going/api/database"
	"github.com/harrisonpim/going/api/models"
)

func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": 200, "message": "👋"})
}

func CreateAuthor(c *fiber.Ctx) error {
	db := database.DB
	author := &models.Author{}
	if err := c.BodyParser(author); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process author"})
	}
	db.Create(&author)
	return c.Status(201).JSON(author)
}

func GetAllAuthors(c *fiber.Ctx) error {
	db := database.DB
	authors := []models.Author{}
	db.Find(&authors)
	return c.JSON(authors)
}

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	author := models.Author{}
	db.Find(&author, id)

	// if author.FirstName == "" {
	// 	return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	// }

	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	author := models.Author{}
	update := new(models.Author)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process author"})
	}
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	}
	// db.Model(&author).Update(&update)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	author := models.Author{}
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No author found with ID: " + id})
	}
	db.Delete(&author)
	return c.JSON(fiber.Map{"status": 200, "message": "Author successfully deleted"})
}

func CreateArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process article"})
	}
	db.Create(&article)
	return c.Status(201).JSON(article)
}

func GetAllArticles(c *fiber.Ctx) error {
	db := database.DB
	articles := []models.Article{}
	db.Find(&articles)
	return c.JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	article := models.Article{}

	db.Find(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	return c.JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	article := models.Article{}

	update := new(models.Article)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Couldn't process article"})
	}
	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	// db.Model(&article).Update(&update)
	return c.JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	article := models.Article{}

	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "No article found with ID: " + id})
	}
	db.Delete(&article)
	return c.JSON(fiber.Map{"status": 200, "message": "Article successfully deleted"})
}

func main() {
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", root)

	authors := app.Group("/authors")
	authors.Post("/", CreateAuthor)
	authors.Get("/", GetAllAuthors)
	authors.Get("/:id", GetAuthor)
	authors.Put("/:id", UpdateAuthor)
	authors.Delete("/:id", DeleteAuthor)

	articles := app.Group("/articles")
	articles.Post("/", CreateArticle)
	articles.Get("/", GetAllArticles)
	articles.Get("/:id", GetArticle)
	articles.Put("/:id", UpdateArticle)
	articles.Delete("/:id", DeleteArticle)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
