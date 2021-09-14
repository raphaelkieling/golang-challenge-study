package main

import (
	"delivery-much-challenge/datasource"
	"delivery-much-challenge/mocks"
	"delivery-much-challenge/service"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Documentation
// How make unit test using fiber, without need run: https://github.com/gofiber/fiber/issues/41
func TestShouldReturnAServer(t *testing.T) {
	server := createServer(ServerConfig{})
	assert.NotEmpty(t, server, "Server is empty")
}

func TestShouldReturnWelcome(t *testing.T) {
	server := createServer(ServerConfig{})

	req, _ := http.NewRequest("GET", "/welcome", nil)
	body, _ := server.Test(req)
	result, _ := io.ReadAll(body.Body)
	assert.Equal(t, string(result), "Hello, World 👋")
	assert.Equal(t, body.StatusCode, 200)
}

func TestShouldReturnNilWhenTakeProducts(t *testing.T) {
	expectedValue := []datasource.Product{}

	mockRepository := &mocks.ProductRepositoryMock{}
	mockRepository.On("GetAllByName", "test").Return(expectedValue, nil)
	productService := service.ProductService{
		ProductRepository: mockRepository,
	}
	server := createServer(ServerConfig{
		productService: &productService,
	})

	req, _ := http.NewRequest("GET", "/products/test", nil)
	body, _ := server.Test(req)
	result, _ := io.ReadAll(body.Body)

	data, _ := json.Marshal(expectedValue)
	assert.Equal(t, string(result), string(data))
	assert.Equal(t, body.StatusCode, 200)
}

func TestShouldSaveANewOrder(t *testing.T) {
	// Mocking
	orderCreatedExpected := datasource.Order{
		Products: []datasource.OrderProduct{
			{
				Name:     "Test-Product-Name",
				Quantity: 2,
				Price:    10,
			},
		},
		Total: 20.0,
	}

	mockOrderRepository := &mocks.OrderRepositoryMock{}
	mockProductRepository := &mocks.ProductRepositoryMock{}

	orderService := service.OrderService{
		ProductRepository: mockProductRepository,
		OrderRepository:   mockOrderRepository,
	}

	mockOrderRepository.On("Save", orderCreatedExpected).Return(orderCreatedExpected, nil)

	mockedProducts := []datasource.Product{
		{
			Name:     "Test-Product-Name",
			Price:    10,
			Quantity: 1,
		},
	}
	mockProductRepository.On("GetAllByName", mockedProducts[0].Name).Return(mockedProducts, nil)

	server := createServer(ServerConfig{
		orderService: &orderService,
	})

	req, _ := http.NewRequest("POST", "/orders", strings.NewReader(`{"products":[{"quantity":2,"name":"Test-Product-Name"}]}`))
	body, _ := server.Test(req)
	result, _ := io.ReadAll(body.Body)

	data, _ := json.Marshal(orderCreatedExpected)
	assert.Equal(t, string(result), string(data))
	assert.Equal(t, body.StatusCode, 200)
}
