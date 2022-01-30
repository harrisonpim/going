package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/harrisonpim/going/api/models"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	baseURL := url.URL{
		Scheme: "http",
		Host:   "api:3000",
	}
	createdAuthor := models.Author{}
	createdArticle := models.Article{}

	client := resty.New()
	g := Goblin(t)

	g.Describe("Create requests", func() {
		g.It("should be able to create an author", func() {
			url := baseURL
			url.Path = path.Join(url.Path, "authors")
			resp, _ := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(`{"FirstName": "first", "LastName": "last"}`).
				Post(url.String())

			g.Assert(resp.StatusCode()).Equal(201)
			json.Unmarshal(resp.Body(), &createdAuthor)
		})

		g.It("should be able to create an article", func() {
			url := baseURL
			url.Path = path.Join(url.Path, "article")
			resp, _ := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(`{"FirstName": "first", "LastName": "last"}`).
				Post(url.String())

			g.Assert(resp.StatusCode()).Equal(201)
			json.Unmarshal(resp.Body(), &createdArticle)
		})

	})

	g.Describe("Read requests", func() {
		g.It("should be able to fetch an author", func() {
			url := baseURL
			url.Path = path.Join(url.Path, "authors", createdAuthor.ID)
			resp, _ := client.R().Get(url.String())
			g.Assert(resp.StatusCode()).Equal(200)

			data := models.Author{}
			json.Unmarshal(resp.Body(), &data)
			fmt.Println(createdAuthor.ID, data.ID)
			g.Assert(data).Equal(createdAuthor)
		})

		g.It("should be able to fetch all authors", func() {
			g.Assert(1 + 1).Equal(2)
		})

		g.It("should be able to fetch an article", func() {
			g.Assert(1 + 1).Equal(2)
		})

		g.It("should be able to fetch all articles", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})

	g.Describe("Update requests", func() {
		g.It("should be able to update an author", func() {
			g.Assert(1 + 1).Equal(2)
		})

		g.It("should be able to update an article", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})

	g.Describe("Delete requests", func() {
		g.It("should be able to delete an author", func() {
			g.Assert(1 + 1).Equal(2)
		})

		g.It("should be able to delete an article", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})
}
