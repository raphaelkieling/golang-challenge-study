package mocks

import (
	"delivery-much-challenge/datasource"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetAllByName(name string) ([]datasource.Product, error) {
	args := m.Called(name)
	return args.Get(0).([]datasource.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Populate(products []datasource.Product) error {
	args := m.Called(products)
	return args.Error(0)
}
