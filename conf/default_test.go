package conf_test

import (
	"delivery-much-challenge/conf"
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
	conf.Env(mockLoader.Load)
	mockLoader.AssertExpectations(t)
	mockLoader.AssertNumberOfCalls(t, "Load", 1)
}

func TestShouldReturnAEmptyConfig(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	env := conf.Env(mockLoader.Load)
	assert.Equal(t, env, &conf.Config{
		Port: "3000",
	})
}

func TestShouldPopulateConfig(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	os.Setenv("APP_PORT", "test")
	env := conf.Env(mockLoader.Load)
	assert.Equal(t, env, &conf.Config{
		Port: "test",
	})
}

func TestShouldHaveDefaultPortValue(t *testing.T) {
	mockLoader := new(mocks.EnvironmentLoaderMock)
	mockLoader.On("Load", mock.Anything).Return(nil)

	os.Clearenv()
	os.Setenv("APP_PORT", "3000")
	env := conf.Env(mockLoader.Load)
	assert.Equal(t, env, &conf.Config{
		Port: "3000",
	})
}
