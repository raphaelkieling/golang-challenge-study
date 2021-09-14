package datasource

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Quantity float64
	Price    float64
	OrderID  uint
}

type Order struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Products []OrderProduct
	Total    float64
}
