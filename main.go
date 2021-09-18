package main

import (
	"delivery-much-challenge/conf"
	"delivery-much-challenge/datasource"
	"delivery-much-challenge/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func initDatabaseAndMigrate(environmentConfiguration *conf.Config) *gorm.DB {
	database := datasource.DatabaseConnect(environmentConfiguration)

	// Inicia os modelos dentro do banco de dados
	database.AutoMigrate(&datasource.Product{})
	database.AutoMigrate(&datasource.OrderProduct{})
	database.AutoMigrate(&datasource.Order{})

	return database
}

func initServer(database *gorm.DB, environmentConfiguration *conf.Config) *fiber.App {
	productReposity := datasource.NewProductRepository(database)
	orderRepository := datasource.NewOrderRepository(database)

	app := createServer(ServerConfig{
		config: environmentConfiguration,
		productService: &service.ProductService{
			ProductRepository: productReposity,
		},
		orderService: &service.OrderService{
			ProductRepository: productReposity,
			OrderRepository:   orderRepository,
		},
	})

	return app
}

func initRabbitMQConsumer(environmentConfiguration *conf.Config) {
	amqpConsumer := AMQPStockConsumer{
		config: environmentConfiguration,
	}

	amqpConsumer.Start()
}

func main() {
	environmentConfiguration := conf.Env(godotenv.Load)

	database := initDatabaseAndMigrate(environmentConfiguration)
	app := initServer(database, environmentConfiguration)
	initRabbitMQConsumer(environmentConfiguration)

	app.Listen(fmt.Sprintf(":%s", environmentConfiguration.Port))
}
