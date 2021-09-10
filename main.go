package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	environmentConfiguration := conf.Env(godotenv.Load)
	database := datasource.DatabaseConnect(environmentConfiguration)

	app := createServer(ServerConfig{
		config:            environmentConfiguration,
		productRepository: datasource.NewProductRepository(database),
	})

	app.Listen(fmt.Sprintf(":%s", environmentConfiguration.Port))
}
