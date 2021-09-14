package conf

import (
	"os"
)

const DEFAULT_PORT = "3000"

type Config struct {
	DbHost     string
	DbPassword string
	DbUser     string
	DbName     string
	DbPort     string
	Port       string
	AmqpHost   string
	AmqpUser   string
	AmqpPass   string
}

func Env(load func(filenames ...string) (err error)) *Config {
	load()

	configuration := &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		Port:       os.Getenv("APP_PORT"),
		AmqpHost:   os.Getenv("AMQP_HOST"),
		AmqpUser:   os.Getenv("AMQP_USER"),
		AmqpPass:   os.Getenv("AMQP_PASSWORD"),
	}

	if configuration.Port == "" {
		configuration.Port = DEFAULT_PORT
	}

	return configuration
}
