package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/service"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	config         *conf.Config
	productService service.IProductService
}

func createServer(conf ServerConfig) *fiber.App {
	app := fiber.New()

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋")
	})

	app.Get("/products/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		products, _ := conf.productService.GetAllByName(name)
		return c.JSON(products)
	})

	return app
}
