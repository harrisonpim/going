# API

A simple CRUD API written with [fiber](https://gofiber.io/), [gorm](https://gorm.io/) and a postgresql database.

## Run

`docker compose up --build api`

## Tests

I've written a suite of tests for the API with [goblin](github.com/franela/goblin) and [resty](github.com/go-resty/resty/v2). You can run them with the command

`docker compose up --build test`
