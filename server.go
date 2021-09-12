package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"

	"github.com/gofiber/fiber/v2"
)

type IProductRepository interface {
	GetAll() ([]datasource.Product, error)
	Populate(products []datasource.Product) error
}

type ServerConfig struct {
	config            *conf.Config
	productRepository IProductRepository
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
