package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/dto"
	"delivery-much-challenge/service"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	config         *conf.Config
	productService service.IProductService
	orderService   service.IOrderService
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

	app.Post("/orders", func(c *fiber.Ctx) error {
		body := c.Body()
		orderCreate, _ := dto.UnmarshalOrderCreate(body)
		orderSaved, _ := conf.orderService.Save(orderCreate)
		return c.JSON(orderSaved)
	})

	return app
}
