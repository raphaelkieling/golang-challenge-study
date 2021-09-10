package mocks

import (
	"delivery-much-challenge/datasource"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetAll() ([]datasource.Product,err error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (p *ProductRepositoryMock) Populate(products []Product) error {
	args := m.Called(products)
	return args.Error(0)
}
