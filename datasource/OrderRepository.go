package datasource

import (
	"gorm.io/gorm"
)

type IOrderRepository interface {
	Save(order Order) (Order, error)
}

type OrderRepository struct {
	connection *gorm.DB
}

// OrderRepository Cria um novo repositório
func NewOrderRepository(connection *gorm.DB) *OrderRepository {
	return &OrderRepository{
		connection: connection,
	}
}

func (p *OrderRepository) Save(order Order) (Order, error) {
	result := p.connection.Create(&order)

	if result.Error != nil {
		return Order{}, result.Error
	}

	return order, nil
}
