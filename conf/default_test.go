package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAEmptyConfig(t *testing.T) {
	os.Clearenv()
	env := Env()
	assert.Equal(t, env, &Config{
		Port: "3000",
	})
}

func TestShouldPopulateConfig(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_PORT", "test")
	env := Env()
	assert.Equal(t, env, &Config{
		Port: "test",
	})
}

func TestShouldHaveDefaultPortValue(t *testing.T) {
	os.Clearenv()
	os.Setenv("APP_PORT", "3000")
	env := Env()
	assert.Equal(t, env, &Config{
		Port: "3000",
	})
}
