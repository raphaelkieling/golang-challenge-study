package service

import (
	"delivery-much-challenge/datasource"
	"delivery-much-challenge/dto"
	"errors"
)

type IOrderService interface {
	Save(orderCreate dto.OrderCreate) (datasource.Order, error)
}

type OrderService struct {
	OrderRepository   datasource.IOrderRepository
	ProductRepository datasource.IProductReposity
}

func (p *OrderService) Save(orderCreate dto.OrderCreate) (datasource.Order, error) {
	orderProducts := []datasource.OrderProduct{}
	sumOfProductPrices := 0.0

	for _, productOrderCreate := range orderCreate.Products {
		productsFoundByName, err := p.ProductRepository.GetAllByName(productOrderCreate.Name)

		if err != nil || len(productsFoundByName) == 0 {
			return datasource.Order{}, errors.New("produto especificado não encontrado")
		}

		product := productsFoundByName[0]
		sumOfProductPrices += product.Price * productOrderCreate.Quantity

		if (productOrderCreate.Quantity - product.Quantity) >= 0 {
			return datasource.Order{}, errors.New("produto fora de estoque")
		}

		orderProducts = append(orderProducts, datasource.OrderProduct{
			Name:     productOrderCreate.Name,
			Price:    product.Price,
			Quantity: productOrderCreate.Quantity,
		})
	}

	order := datasource.Order{
		Products: orderProducts,
		Total:    sumOfProductPrices,
	}

	return p.OrderRepository.Save(order)
}