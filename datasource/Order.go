package datasource

type OrderProduct struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Quantity float64
	Price    float64
}

type Order struct {
	ID       uint `gorm:"primaryKey"`
	Products []OrderProduct
	Total    float64
}
