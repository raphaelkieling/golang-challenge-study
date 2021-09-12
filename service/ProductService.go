package service

import "delivery-much-challenge/datasource"

type IProductService interface {
	GetAllByName(name string) ([]datasource.Product, error)
}

type ProductService struct {
	ProductRepository datasource.IProductReposity
}

func (p *ProductService) GetAllByName(name string) ([]datasource.Product, error) {
	return p.ProductRepository.GetAllByName(name)
}
