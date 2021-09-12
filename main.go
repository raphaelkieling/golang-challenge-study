package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"
	"delivery-much-challenge/service"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	environmentConfiguration := conf.Env(godotenv.Load)
	database := datasource.DatabaseConnect(environmentConfiguration)

	// Inicia os modelos dentro do banco de dados
	// TODO: Posso colocar em outro arquivo datasource.MigrateModels() ou algo do tipo
	database.AutoMigrate(&datasource.Product{})
	database.AutoMigrate(&datasource.OrderProduct{})
	database.AutoMigrate(&datasource.Order{})

	app := createServer(ServerConfig{
		config: environmentConfiguration,
		productService: &service.ProductService{
			ProductRepository: datasource.NewProductRepository(database),
		},
	})

	app.Listen(fmt.Sprintf(":%s", environmentConfiguration.Port))
}
