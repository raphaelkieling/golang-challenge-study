package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	config            *conf.Config
	productRepository *datasource.ProductRepository
}

func createServer(conf ServerConfig) *fiber.App {
	app := fiber.New()

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		products, _ := conf.productRepository.GetAll()
		return c.JSON(products)
	})

	return app
}
