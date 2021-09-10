package main

import (
	"delivery-much-challenge/mocks"
	"io"
	"net/http"
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
	mockRepository := mocks.ProductRepositoryMock{}
	server := createServer(ServerConfig{
		productRepository: mockRepository,
	})

	req, _ := http.NewRequest("GET", "/products", nil)
	body, _ := server.Test(req)
	result, _ := io.ReadAll(body.Body)
	assert.Equal(t, string(result), "Hello, World 👋")
	assert.Equal(t, body.StatusCode, 200)
}
