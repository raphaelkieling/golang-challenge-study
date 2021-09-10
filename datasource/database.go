package datasource

import (
	"delivery-much-challenge/conf"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect(conf *conf.Config) *gorm.DB {
	fmt.Println(conf)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.DbHost, conf.DbUser, conf.DbPassword, conf.DbName, conf.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected")
	return db
}
