package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig("test", "localhost", "8080")

	assert.Equal(t, config.AppName, "test")
	assert.Equal(t, config.Host, "localhost")
	assert.Equal(t, config.Port, "8080")
}
