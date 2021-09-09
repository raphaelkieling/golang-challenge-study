package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "3000"

type Config struct {
	DbHost     string
	DbPassword string
	DbUser     string
	DbName     string
	DbPort     string
	Port       string
}

func Env() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, trying get from env")
	}

	conf := &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		Port:       os.Getenv("APP_PORT"),
	}

	fmt.Println(conf)

	if conf.Port == "" {
		conf.Port = DEFAULT_PORT
	}

	return conf
}
