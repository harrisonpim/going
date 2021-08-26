package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisonpim/going/api/database"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateAuthor(c *fiber.Ctx) error {
	db := database.DBConn
	author := new(database.Author)
	if err := c.BodyParser(author); err != nil {
		return c.Status(400).SendString("Couldn't process author")
	}
	db.Create(&author)
	return c.JSON(author)
}

func AllAuthors(c *fiber.Ctx) error {
	db := database.DBConn
	var authors []database.Author
	db.Find(&authors)
	return c.JSON(authors)
}

func SingleAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var author database.Author
	db.Find(&author, id)
	if author.FirstName == "" {
		return c.Status(404).SendString("No author found with ID: " + id)
	}
	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var author database.Author
	update := new(database.Author)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).SendString("Couldn't process author")
	}
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).SendString("No author found with ID: " + id)
	}
	db.Model(&author).Update(&update)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var author database.Author
	db.First(&author, id)
	if author.FirstName == "" {
		return c.Status(404).SendString("No author found with ID: " + id)
	}
	db.Delete(&author)
	return c.SendString("Author successfully deleted")
}

func CreateArticle(c *fiber.Ctx) error {
	db := database.DBConn
	article := new(database.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(400).SendString("Couldn't process article")
	}
	db.Create(&article)
	return c.JSON(article)
}

func AllArticles(c *fiber.Ctx) error {
	db := database.DBConn
	var articles []database.Article
	db.Find(&articles)
	return c.JSON(articles)
}

func SingleArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var article database.Article
	db.Find(&article, id)
	if article.Title == "" {
		return c.Status(404).SendString("No article found with ID: " + id)
	}
	return c.JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var article database.Article
	update := new(database.Article)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).SendString("Couldn't process article")
	}
	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).SendString("No article found with ID: " + id)
	}
	db.Model(&article).Update(&update)
	return c.JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var article database.Article
	db.First(&article, id)
	if article.Title == "" {
		return c.Status(404).SendString("No article found with ID: " + id)
	}
	db.Delete(&article)
	return c.SendString("Article successfully deleted")
}

func main() {
	database.Init()

	app := fiber.New()

	app.Post("/articles", CreateArticle)
	app.Get("/articles", AllArticles)
	app.Get("/articles/:id", SingleArticle)
	app.Put("/articles/:id", UpdateArticle)
	app.Delete("/articles/:id", DeleteArticle)

	app.Post("/authors", CreateAuthor)
	app.Get("/authors", AllAuthors)
	app.Get("/authors/:id", SingleAuthor)
	app.Put("/authors/:id", UpdateAuthor)
	app.Delete("/authors/:id", DeleteAuthor)

	app.Listen(":3000")
}
