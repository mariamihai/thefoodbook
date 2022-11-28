package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalConfiguration(t *testing.T) {
	config := InitializeConfiguration()

	assert.Equal(t, 3000, config.Port, "Configured port should be 3000.")

	//assert.NotEmpty(t, config.Uri)
	//assert.NotEmpty(t, config.Username)
	//assert.NotEmpty(t, config.Password)
}
