package main

import (
	"delivery-much-challenge/conf"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func createServer(conf *conf.Config) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, World 👋! %s", conf.DbName))
	})

	return app
}
