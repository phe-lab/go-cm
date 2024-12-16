package test

import (
	"os"
	"testing"

	"github.com/phe-lab/go-cm/cfg"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_FromDefaultValues(t *testing.T) {
	// Declare default values
	defaultValues := map[string]interface{}{
		"env":         "DefaultApp",
		"server.port": 3000,
	}

	// Setup OS environment variables
	os.Setenv("APP_ENV", "EnvApp")
	os.Setenv("APP_SERVER_HOST", "0.0.0.0")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("APP_SERVER_HOST")

	// Checking
	var config Config
	err := cfg.LoadConfig(&config, cfg.LoadOptions{
		EnvPrefix:     "APP_",
		DefaultValues: &defaultValues,
	})
	assert.NoError(t, err)
	assert.Equal(t, "EnvApp", config.Env)
	assert.Equal(t, 3000, config.Server.Port)
	assert.Equal(t, "0.0.0.0", config.Server.Host)
}
