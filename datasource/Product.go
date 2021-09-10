package datasource

type Product struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Price    float64
	Quantity float64
}
