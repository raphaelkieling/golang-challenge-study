package conf

import (
	"delivery-much-challenge/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldTryPopulateUsingTheLoader(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	Env(mockLoader.Load)
	mockLoader.AssertExpectations(t)
	mockLoader.AssertNumberOfCalls(t, "Load", 1)
}

func TestShouldReturnAEmptyConfig(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	env := Env(mockLoader.Load)
	assert.Equal(t, env, &Config{
		Port: "3000",
	})
}

func TestShouldPopulateConfig(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	os.Setenv("APP_PORT", "test")
	env := Env(mockLoader.Load)
	assert.Equal(t, env, &Config{
		Port: "test",
	})
}

func TestShouldHaveDefaultPortValue(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	os.Setenv("APP_PORT", "3000")
	env := Env(mockLoader.Load)
	assert.Equal(t, env, &Config{
		Port: "3000",
	})
}
