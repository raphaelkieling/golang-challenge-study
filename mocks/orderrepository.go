package mocks

import (
	"delivery-much-challenge/datasource"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) Save(order datasource.Order) (datasource.Order, error) {
	args := m.Called(order)
	return args.Get(0).(datasource.Order), args.Error(1)
}
