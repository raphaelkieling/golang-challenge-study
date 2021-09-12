package datasource

import (
	"gorm.io/gorm"
)

type IProductReposity interface {
	GetAllByName(name string) ([]Product, error)
}

type ProductRepository struct {
	connection *gorm.DB
}

// NewProductRepository Cria um novo repositório
func NewProductRepository(connection *gorm.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetAllByName(name string) ([]Product, error) {
	var products []Product
	result := p.connection.Where("Name = ?", name).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductRepository) Populate(products []Product) error {
	for i := 0; i < len(products); i++ {
		product := products[i]
		result := p.connection.Create(&product)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
